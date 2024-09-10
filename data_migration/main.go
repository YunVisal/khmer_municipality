package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	out, err := os.Create("../data/data.xlsx")
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer out.Close()

	req, err := http.NewRequest("GET", "https://content.dropboxapi.com/2/files/download", nil)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	req.Header.Set("Dropbox-API-Arg", "{\"path\":\"id:qnLpP3P8qUAAAAAAAAAAHw\"}")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("bad status: %s", resp.Status)
		return
	}

	// Copy the response body into the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return
	}

	// xl, err := xlsxreader.OpenFile("../data/ncdd_admin_database_25provinces__2023.xlsx")
	// if err != nil {
	// 	fmt.Println("ERROR!")
	// 	fmt.Println(err)
	// 	return
	// }
	// defer xl.Close()

	// CreateProvinceCSVFile(xl.Sheets)
	// CreateDistrictCSVFile(xl)
	// CreateCommuneCSVFile(xl)
	// CreateVillageCSVFile(xl)
}
