package daytime

import (
	"bytes"
	"fmt"
	"strconv"
)

// Seconds since "noon minus 12h"
// (effectively midnight, except for days on which daylight savings time changes occur).
type Daytime int

func (d *Daytime) UnmarshalText(text []byte) error {
	hms := bytes.SplitN(text, []byte{58}, 3)
	ss, err := strconv.Atoi(string(hms[2]))
	if err != nil {
		return err
	}
	mm, err := strconv.Atoi(string(hms[1]))
	if err != nil {
		return err
	}
	hh, err := strconv.Atoi(string(hms[0]))
	if err != nil {
		return err
	}
	*d = Daytime(ss + mm*60 + hh*3600)
	return nil
}

/*
func (d Daytime) MarshalText() ([]byte, error) {
	hh := int(d / 3600)
	mm := int((d % 3600) / 60)
	ss := d % 60

	return []byte(fmt.Sprintf("%02f:%02f:%02f", hh, mm, ss)), nil
}
*/

func (d Daytime) String() string {
	hh := int(int(d) / 3600)
	mm := int((int(d) % 3600) / 60)
	ss := int(d) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
}
