package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/thedatashed/xlsxreader"
)

func CreateDistrictCSVFile(rows chan xlsxreader.Row, file *os.File) {
	districts := [][]string{}
	i := 0
	for row := range rows {
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

	w := csv.NewWriter(file)
	w.WriteAll(districts)
	if err := w.Error(); err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
}
