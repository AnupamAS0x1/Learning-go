package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("connected with forntend")
	PerformGetRequest()
}

func PerformGetRequest() {

	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content ", response.ContentLength)

	var responseStrings strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseStrings.Write(content)

	fmt.Println("Bytecount is:", byteCount)
	fmt.Println(responseStrings.String())

	fmt.Println(string(content))
}
