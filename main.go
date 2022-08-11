package main  

import ( 
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://api.polygon.io/v1/open-close/AAPL/2020-10-14?adjusted=true&apiKey=3WhkjwKYZY4V6wK0su0xzcPFSN0FlzvU")
	if err != nil {
		fmt.Println(err)
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} 
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}