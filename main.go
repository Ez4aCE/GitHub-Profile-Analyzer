package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.github.com/users/Ez4aCE"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Error fetching data : ", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response : ", err)
	}

	fmt.Println(string(body))
}
