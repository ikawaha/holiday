package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ikawaha/holiday/download"
)

const (
	holidayData = "./holidays.json"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	h, err := download.Holidays(context.Background())
	if err != nil {
		return err
	}
	newer, err := json.MarshalIndent(h, "", "\t")
	if err != nil {
		return err
	}
	old, err := os.ReadFile(holidayData)
	if err != nil {
		return err
	}
	if !bytes.Equal(newer, old) {
		if err := os.WriteFile(holidayData, newer, 0600); err != nil {
			return err
		}
	}
	return nil
}
