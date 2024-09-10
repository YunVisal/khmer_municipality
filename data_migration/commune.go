package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/thedatashed/xlsxreader"
)

func CreateCommuneCSVFile(xl *xlsxreader.XlsxFileCloser) {
	COMMUNE_FILE_DIR := "../data/commune.csv"
	file, err := os.Create(COMMUNE_FILE_DIR)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer file.Close()

	communes := [][]string{
		{"id", "name", "english_name", "district_id"},
	}
	for _, sheet := range xl.Sheets {
		i := 0
		for row := range xl.ReadRows(sheet) {
			if i < 2 {
				i++
				continue
			}
			if row.Cells[0].Value == "ឃុំ" || row.Cells[0].Value == "សង្កាត់" {
				id := row.Cells[1].Value
				name := row.Cells[2].Value
				english_name := row.Cells[3].Value
				district_id := row.Cells[1].Value[:len(row.Cells[1].Value)-2]
				commune := []string{id, name, english_name, district_id}
				communes = append(communes, commune)
			}
		}
	}
	w := csv.NewWriter(file)
	w.WriteAll(communes)
	if err := w.Error(); err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
}
