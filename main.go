package main  

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type Stock struct {
	Status 		string 
	From 		string
	Symbol 		string
	Open 		float64
	High 		float64
	Low 		float64
	Close 		float64
	Volume 		int
	AfterHours 	int
	PreMarket 	float64
}

var voo_value string

func server(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "VOO at close: " + voo_value)
    fmt.Println("Data accessed")
}

func handleRequests() {
    http.HandleFunc("/", server)
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	voo_value = get_stock_value(voo_open_close)
	fmt.Println("Value of VOO at close: " + voo_value)
	handleRequests()
}

func get_stock_value(endpoint string) string {
	resp, err := http.Get(endpoint)
	if err != nil {
		return err.Error()
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err.Error()
		} else{
			var response Stock
			json.Unmarshal(bodyBytes, &response)
			return strconv.FormatFloat(response.Close, 'f', 2, 64)
		}
	}
}