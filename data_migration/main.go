package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// func main() {
// 	out, err := os.Create("../data/data.xlsx")
// 	if err != nil {
// 		fmt.Println("ERROR!")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer out.Close()

// 	req, err := http.NewRequest("GET", "https://content.dropboxapi.com/2/files/download", nil)
// 	if err != nil {
// 		fmt.Println("ERROR!")
// 		fmt.Println(err)
// 		return
// 	}
// 	req.Header.Set("Dropbox-API-Arg", "{\"path\":\"id:qnLpP3P8qUAAAAAAAAAAHw\"}")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("ERROR!")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Check if the request was successful
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Errorf("bad status: %s", resp.Status)
// 		return
// 	}

// 	// Copy the response body into the file
// 	_, err = io.Copy(out, resp.Body)
// 	if err != nil {
// 		fmt.Println("ERROR!")
// 		fmt.Println(err)
// 		return
// 	}

// 	// xl, err := xlsxreader.OpenFile("../data/ncdd_admin_database_25provinces__2023.xlsx")
// 	// if err != nil {
// 	// 	fmt.Println("ERROR!")
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// defer xl.Close()

// 	// CreateProvinceCSVFile(xl.Sheets)
// 	// CreateDistrictCSVFile(xl)
// 	// CreateCommuneCSVFile(xl)
// 	// CreateVillageCSVFile(xl)
// }

func main() {
	http.HandleFunc("/", servePage)

	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	page_path := ""
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		page_path = "./static/index.html"
	} else if r.URL.Path == "/token/" {
		page_path = "./static/token_redirect.html"
	}

	if page_path != "" {
		tmpl, err := template.ParseFiles(page_path)
		if err != nil {
			http.Error(w, "Error loading form", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}
