package main  

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

type Response struct {
	status 		string 
	from 		string
	symbol 		string
	open 		string
	high 		string
	low 		string
	close 		string
	volume 		string
	afterHours 	string
	preMarket 	string
}

func main() {
	fmt.Println(get_stock_value(voo_open_close))
	time.Sleep(5 * time.Second)
	//fmt.Println(get_stock_value(voo_open_close))
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
			//bodyString := string(bodyBytes)
			var response Response
			json.Unmarshal(bodyBytes, &response)
			return response.close
		}
	}
}