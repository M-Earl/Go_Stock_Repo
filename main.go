package main  

import ( 
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println(get_stock_value(voo_open_close))
	time.Sleep(5 * time.Second)
	fmt.Println(get_stock_value(voo_open_close))
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
			bodyString := string(bodyBytes)
			return bodyString
		}
	}
}