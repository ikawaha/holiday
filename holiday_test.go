package holiday

import (
	"fmt"
	"sort"
	"strings"
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

func ExampleIsHolidayYYYYMMDD() {
	testdata := []string{"2021/07/23", "2021/10/10"}
	for _, v := range testdata {
		day, ok := IsHolidayYYYYMMDD(v)
		fmt.Printf("IsHolidayYYYYMMDD(%s)= %q, %v\n", v, day, ok)
	}
	// Output:
	// IsHolidayYYYYMMDD(2021/07/23)= "スポーツの日", true
	// IsHolidayYYYYMMDD(2021/10/10)= "", false
}

func ExampleIsHoliday() {
	testdata := []string{"2021/07/23", "2021/10/10"}
	for _, v := range testdata {
		t, err := time.Parse(DateFormat, v)
		if err != nil {
			panic(err)
		}
		day, ok := IsHoliday(t)
		fmt.Printf("IsHoliday(%s)= %q, %v\n", v, day, ok)
	}
	// Output:
	// IsHoliday(2021/07/23)= "スポーツの日", true
	// IsHoliday(2021/10/10)= "", false
}

func ExampleCatalog() {
	var year2021 []string
	for k := range Catalog {
		if strings.HasPrefix(k, "2021/") {
			year2021 = append(year2021, k)
		}
	}
	sort.Strings(year2021)
	for _, v := range year2021 {
		fmt.Println(v, Catalog[v])
	}
	// Output:
	// 2021/01/01 元日
	// 2021/01/11 成人の日
	// 2021/02/11 建国記念の日
	// 2021/02/23 天皇誕生日
	// 2021/03/20 春分の日
	// 2021/04/29 昭和の日
	// 2021/05/03 憲法記念日
	// 2021/05/04 みどりの日
	// 2021/05/05 こどもの日
	// 2021/07/22 海の日
	// 2021/07/23 スポーツの日
	// 2021/08/08 山の日
	// 2021/08/09 休日
	// 2021/09/20 敬老の日
	// 2021/09/23 秋分の日
	// 2021/11/03 文化の日
	// 2021/11/23 勤労感謝の日
}
