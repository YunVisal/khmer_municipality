package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/thedatashed/xlsxreader"
)

func CreateCommuneCSVFile(rows chan xlsxreader.Row, file *os.File) {
	communes := [][]string{}
	i := 0
	for row := range rows {
		if i < 2 {
			i++
			continue
		}
		if row.Cells[0].Value == "ឃុំ" || row.Cells[0].Value == "សង្កាត់" {
			id := row.Cells[1].Value
			name := row.Cells[2].Value
			english_name := row.Cells[3].Value
			province_id := row.Cells[1].Value[:len(row.Cells[1].Value)-2]
			commune := []string{id, name, english_name, province_id}
			communes = append(communes, commune)
		}
	}

	w := csv.NewWriter(file)
	w.WriteAll(communes)
	if err := w.Error(); err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
}
