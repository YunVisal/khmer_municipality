package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/thedatashed/xlsxreader"
)

func CreateVillageCSVFile(rows chan xlsxreader.Row, file *os.File) {
	villages := [][]string{}
	i := 0
	for row := range rows {
		if i < 2 {
			i++
			continue
		}
		if row.Cells[0].Value == "ភូមិ" {
			id := row.Cells[1].Value
			name := row.Cells[2].Value
			english_name := row.Cells[3].Value
			province_id := row.Cells[1].Value[:len(row.Cells[1].Value)-2]
			village := []string{id, name, english_name, province_id}
			villages = append(villages, village)
		}
	}

	w := csv.NewWriter(file)
	w.WriteAll(villages)
	if err := w.Error(); err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
}
