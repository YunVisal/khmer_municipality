package main

import (
	"fmt"
	"os"

	"github.com/thedatashed/xlsxreader"
)

func main() {
	xl, err := xlsxreader.OpenFile("../data/ncdd_admin_database_25provinces__2023.xlsx")
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}

	defer xl.Close()

	CreateProvinceCSVFile(xl.Sheets)

	DISTRICT_FILE_DIR := "../data/district.csv"
	district_file, err := os.Create(DISTRICT_FILE_DIR)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer district_file.Close()

	COMMUNE_FILE_DIR := "../data/commune.csv"
	commune_file, err := os.Create(COMMUNE_FILE_DIR)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer district_file.Close()

	VILLAGE_FILE_DIR := "../data/village.csv"
	village_file, err := os.Create(VILLAGE_FILE_DIR)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer district_file.Close()

	for _, sheet := range xl.Sheets {
		CreateDistrictCSVFile(xl.ReadRows(sheet), district_file)
		CreateCommuneCSVFile(xl.ReadRows(sheet), commune_file)
		CreateVillageCSVFile(xl.ReadRows(sheet), village_file)
	}
}
