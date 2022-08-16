package main  

import ( 
	"fmt"
	"net/http"
	"io/ioutil"
)

const aapl_open_close = "https://api.polygon.io/v1/open-close/AAPL/2020-10-14?adjusted=true&apiKey=3WhkjwKYZY4V6wK0su0xzcPFSN0FlzvU"
const voo_open_close = "https://api.polygon.io/v1/open-close/VOO/2020-10-14?adjusted=true&apiKey=3WhkjwKYZY4V6wK0su0xzcPFSN0FlzvU"
const msft_open_close = "https://api.polygon.io/v1/open-close/MSFT/2020-10-14?adjusted=true&apiKey=3WhkjwKYZY4V6wK0su0xzcPFSN0FlzvU"


func main() {
	resp, err := http.Get(voo_open_close)
	if err != nil {
		fmt.Println(err)
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else{
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		}
	}
}