package download

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ikawaha/holiday"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// see. https://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
const holidayJPURI = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

// HolidayCSV downloads the csv file of the national holidays from a Cabinet Office page.
func HolidayCSV() ([][]string, error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(holidayJPURI) //nolint:noctx
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(transform.NewReader(bytes.NewReader(body), japanese.ShiftJIS.NewDecoder()))
	return r.ReadAll()
}

const csvDateFormat = `2006/1/2`

// Holidays downloads the national holidays from a Cabinet Office page and return it.
func Holidays() (holiday.Holidays, error) {
	records, err := HolidayCSV()
	if err != nil {
		return nil, err
	}
	ret := holiday.Holidays{}
	for i, v := range records {
		if i == 0 {
			continue // skip header: "国民の祝日・休日月日", "国民の祝日・休日名称"
		}
		if len(v) != 2 {
			return nil, fmt.Errorf("unexpected format (expected col size 2), %+v", v)
		}
		t, err := time.Parse(csvDateFormat, v[0])
		if err != nil {
			return nil, fmt.Errorf("unexpected format (expected YYYY/M/D), %s", v[0])
		}
		ret[t.Format(holiday.DateFormat)] = v[1]
	}
	return ret, nil
}
