package sfomuseum

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"strconv"
	"strings"
	"sync"
)

type Airline struct {
	WOFID        int64  `json:"wof:id"`
	Name         string `json:"wof:name"`
	SFOMuseumID  int    `json:"sfomuseum:airline_id"`
	IATACode     string `json:"iata:code,omitempty"`
	ICAOCode     string `json:"icao:code,omitempty"`
	ICAOCallsign string `json:"icao:callsign,omitempty"`
	WikidataID   string `json:"wd:id,omitempty"`
}

func (a *Airline) String() string {
	return fmt.Sprintf("%s %s %s \"%s\" %d", a.IATACode, a.ICAOCode, a.ICAOCallsign, a.Name, a.WOFID)
}

var lookup_table *sync.Map
var lookup_init sync.Once

type SFOMuseumLookup struct {
	airlines.Lookup
}

func NewLookup() (airlines.Lookup, error) {

	var lookup_err error

	lookup_func := func() {

		var airline []*Airline

		err := json.Unmarshal([]byte(AirlineData), &airline)

		if err != nil {
			lookup_err = err
			return
		}

		table := new(sync.Map)

		for idx, craft := range airline {

			pointer := fmt.Sprintf("pointer:%d", idx)
			table.Store(pointer, craft)

			str_wofid := strconv.FormatInt(craft.WOFID, 10)

			possible_codes := []string{
				craft.IATACode,
				craft.ICAOCode,
				craft.ICAOCallsign,
				str_wofid,
			}

			for _, code := range possible_codes {

				if code == "" {
					continue
				}

				pointers := make([]string, 0)
				has_pointer := false

				others, ok := table.Load(code)

				if ok {

					pointers = others.([]string)
				}

				for _, dupe := range pointers {

					if dupe == pointer {
						has_pointer = true
						break
					}
				}

				if has_pointer {
					continue
				}

				pointers = append(pointers, pointer)
				table.Store(code, pointers)
			}

			idx += 1
		}

		lookup_table = table
	}

	lookup_init.Do(lookup_func)

	if lookup_err != nil {
		return nil, lookup_err
	}

	l := SFOMuseumLookup{}
	return &l, nil
}

func (l *SFOMuseumLookup) Find(code string) ([]interface{}, error) {

	pointers, ok := lookup_table.Load(code)

	if !ok {
		return nil, errors.New("Not found")
	}

	airline := make([]interface{}, 0)

	for _, p := range pointers.([]string) {

		if !strings.HasPrefix(p, "pointer:") {
			return nil, errors.New("Invalid pointer")
		}

		row, ok := lookup_table.Load(p)

		if !ok {
			return nil, errors.New("Invalid pointer")
		}

		airline = append(airline, row.(*Airline))
	}

	return airline, nil
}
