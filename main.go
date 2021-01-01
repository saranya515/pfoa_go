package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/csv"
	"strings"
)

type Profile struct {
	Category string `json:"category"`
	Expenditure   int `json:"expenditure"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	profile := []Profile{{"groceries", 450}, {"Sub Total", 1000}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func readCsv() {  
	data, err := ioutil.ReadFile("statement.CSV")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
	in := string(data)
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}

func main() {
	//handleRequests()
	readCsv()
}

