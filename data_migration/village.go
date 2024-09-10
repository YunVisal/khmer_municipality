package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/thedatashed/xlsxreader"
)

func CreateVillageCSVFile(xl *xlsxreader.XlsxFileCloser) {
	VILLAGE_FILE_DIR := "../data/village.csv"
	file, err := os.Create(VILLAGE_FILE_DIR)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer file.Close()

	villages := [][]string{
		{"id", "name", "english_name", "commune_id"},
	}
	for _, sheet := range xl.Sheets {
		i := 0
		for row := range xl.ReadRows(sheet) {
			if i < 2 {
				i++
				continue
			}
			if row.Cells[0].Value == "ភូមិ" {
				id := row.Cells[1].Value
				name := row.Cells[2].Value
				english_name := row.Cells[3].Value
				commune_id := row.Cells[1].Value[:len(row.Cells[1].Value)-2]
				village := []string{id, name, english_name, commune_id}
				villages = append(villages, village)
			}
		}
	}
	w := csv.NewWriter(file)
	w.WriteAll(villages)
	if err := w.Error(); err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
}
