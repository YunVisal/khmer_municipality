package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateProvinceCSVFile(sheet_name_arr []string) {
	XLSX_SHEET_NAME_SEP := ". "
	provinces := [][]string{
		{"id", "name", "english_name"},
	}
	for i, sheet_name := range sheet_name_arr {
		province_name := strings.Split(sheet_name, XLSX_SHEET_NAME_SEP)[1]
		province := []string{strconv.Itoa(i + 1), "", province_name}
		provinces = append(provinces, province)
	}

	PROVINCES_FILE_DIR := "../data/provinces.csv"
	_, err := os.Stat(PROVINCES_FILE_DIR)
	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(PROVINCES_FILE_DIR)
		if err != nil {
			fmt.Println("ERROR!")
			fmt.Println(err)
			return
		}
		defer file.Close()

		w := csv.NewWriter(file)
		w.WriteAll(provinces)
		if err := w.Error(); err != nil {
			fmt.Println("ERROR!")
			fmt.Println(err)
		}
	} else if err == nil {
		fmt.Println("Exist")
	} else {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
}
