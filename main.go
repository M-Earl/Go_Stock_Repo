package main  

import ( 
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

const open_close_endpoint = "https://api.polygon.io/v1/open-close/"
const api_key = "3WhkjwKYZY4V6wK0su0xzcPFSN0FlzvU"
const date = "2020-10-14"
const extra = "?adjusted=true&apiKey="
const aapl = "AAPL/"
const voo = "VOO/"
const msft = "MSFT/"

const aapl_open_close = open_close_endpoint + aapl + date + extra + api_key
const voo_open_close = open_close_endpoint + voo + date + extra + api_key
const msft_open_close = open_close_endpoint + msft + date + extra + api_key

func main() {
	fmt.Println(get_voo())
	time.Sleep(5 * time.Second)
	fmt.Println(get_voo())
}

func get_voo() string {
	resp, err := http.Get(voo_open_close)
	if err != nil {
		return err.Error()
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err.Error()
		} else{
			bodyString := string(bodyBytes)
			return bodyString
		}
	}
}