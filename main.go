package main  

import ( 
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

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