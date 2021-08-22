package holiday

import (
	"testing"
	"time"
)

func Test_CatalogData(t *testing.T) {
	if got, want := len(Catalog), 975; got < want { // Number of days as of 2021.
		t.Fatalf("got %d, want >%d", got, want)
	}
}

func TestIsHolidayGolden(t *testing.T) {
	for date := range Catalog {
		tm, err := time.Parse(DateFormat, date)
		if err != nil {
			t.Errorf("unexpected error, invalid date format: %s", date)
		}
		if _, got := IsHoliday(tm); got == false {
			t.Errorf("IsHoliday(%q) = %v, want true", tm.Format(DateFormat), got)
		}
	}
}

func TestIsHolidayYYYYMMDDGolden(t *testing.T) {
	for date := range Catalog {
		if _, got := IsHolidayYYYYMMDD(date); got == false {
			t.Errorf("IsHolidayYYYYMMDD(%q) = %v, want true", date, got)
		}
	}
}
