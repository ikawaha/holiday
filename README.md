[![Go Reference](https://pkg.go.dev/badge/github.com/ikawaha/holiday.svg)](https://pkg.go.dev/github.com/ikawaha/holiday)

# holiday

This is a list of holidays announced by the Cabinet Office as national holidays. This list
is compared with the Cabinet Office's list once a month, and a pull-request is created when
there is an update.

```go
package main

import (
	"fmt"
	"time"

	"github.com/ikawaha/holiday"
)

func main() {
	dates := []string{"2021/07/23", "2021/10/10"}
	for _, v := range dates {
		// YYYY/MM/DD
		day, ok := holiday.IsHolidayYYYYMMDD(v)
		fmt.Printf("IsHolidayYYYYMMDD(%s)= %q, %v\n", v, day, ok)

		// time.Time
		t, err := time.Parse(holiday.DateFormat, v)
		if err != nil {
			panic(err)
		}
		day, ok = holiday.IsHoliday(t)
		fmt.Printf("IsHoliday(%s)= %q, %v\n", v, day, ok)
	}
}
```

Output:
```text
IsHolidayYYYYMMDD(2021/07/23)= "スポーツの日" true
IsHoliday(2021/07/23)= "スポーツの日" true
IsHolidayYYYYMMDD(2021/10/10)= "" false
IsHoliday(2021/10/10)= "" false
```

---
MIT

