package main

import (
	"fmt"

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
	CreateDistrictCSVFile(xl)
	CreateCommuneCSVFile(xl)
	CreateVillageCSVFile(xl)
}
