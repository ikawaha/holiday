package holiday

import (
	_ "embed"
	"encoding/json"
	"time"
)

type Holidays map[string]string

var (
	// Catalog is the list of National Holidays.
	Catalog Holidays

	//go:embed holidays.json
	holidaysJSON []byte
)

const DateFormat = `2006/01/02`

func init() {
	Catalog = func() Holidays {
		ret := map[string]string{}
		if err := json.Unmarshal(holidaysJSON, &ret); err != nil {
			panic(err)
		}
		return ret
	}()
}

// IsHoliday returns the name of a holiday and true if the given time is a national holiday.
func IsHoliday(t time.Time) (string, bool) {
	return IsHolidayYYYYMMDD(t.Format(DateFormat))
}

// IsHolidayYYYYMMDD returns the name of a holiday and true if the given date
// in YYYY/MM/DD format (eg. 2006/01/02) is a national holiday.
func IsHolidayYYYYMMDD(date string) (string, bool) {
	name, ok := Catalog[date]
	return name, ok
}
