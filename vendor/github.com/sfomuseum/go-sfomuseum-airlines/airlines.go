package airlines

// I do not look this, returning []interface{} instead of something
// less-obtuse but there isn't really any commonality (yet...) between
// the Airline thingies defined in the wikipedia/sfomuseum packages...
// (20190430/thisisaaronland)

type Lookup interface {
	Find(string) ([]interface{}, error)
}

/*
package airlines

import (
	"fmt"
)

type Airline struct {
	IATA      string // 2-letter code
	ICAO      string // 3-letter code
	TELEPHONY string
	Name      string
	Duplicate bool // IATA, this should probably be renamed
}

func (a *Airline) String() string {
	return fmt.Sprintf("%s %s %s", a.IATA, a.ICAO, a.TELEPHONY)
}
*/
