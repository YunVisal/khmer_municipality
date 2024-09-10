package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/thedatashed/xlsxreader"
)

func CreateDistrictCSVFile(xl *xlsxreader.XlsxFileCloser) {
	DISTRICT_FILE_DIR := "../data/district.csv"
	file, err := os.Create(DISTRICT_FILE_DIR)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer file.Close()

	districts := [][]string{
		{"id", "name", "english_name", "province_id"},
	}
	for _, sheet := range xl.Sheets {
		i := 0
		for row := range xl.ReadRows(sheet) {
			if i < 2 {
				i++
				continue
			}
			if row.Cells[0].Value == "ស្រុក" || row.Cells[0].Value == "ក្រុង" {
				id := row.Cells[1].Value
				name := row.Cells[2].Value
				english_name := row.Cells[3].Value
				province_id := row.Cells[1].Value[:len(row.Cells[1].Value)-2]
				district := []string{id, name, english_name, province_id}
				districts = append(districts, district)
			}
		}
	}
	w := csv.NewWriter(file)
	w.WriteAll(districts)
	if err := w.Error(); err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
}
