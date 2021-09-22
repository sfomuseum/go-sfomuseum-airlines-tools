package sfomuseum

import (
	sfomuseum_props "github.com/sfomuseum/go-sfomuseum-geojson/properties/sfomuseum"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
)

func SFOMuseumAirlineFromFeature(f geojson.Feature) (*Airline, error) {

	pt := sfomuseum_props.Placetype(f)

	if pt != "airline" {
		return nil, nil
	}

	wof_id := whosonfirst.Id(f)
	name := whosonfirst.Name(f)

	sfom_id := utils.Int64Property(f.Bytes(), []string{"properties.sfomuseum:airline_id"}, -1)

	concordances, err := whosonfirst.Concordances(f)

	if err != nil {
		return nil, err
	}

	a := &Airline{
		WOFID:       wof_id,
		SFOMuseumID: int(sfom_id),
		Name:        name,
	}

	iata_code, ok := concordances["iata:code"]

	if ok {
		a.IATACode = iata_code
	}

	icao_code, ok := concordances["icao:code"]

	if ok {
		a.ICAOCode = icao_code
	}

	callsign, ok := concordances["icao:callsign"]

	if ok {
		a.ICAOCallsign = callsign
	}

	id, ok := concordances["wd:id"]

	if ok {
		a.WikidataID = id
	}

	return a, nil
}
