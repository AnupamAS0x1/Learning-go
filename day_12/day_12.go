package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("plyaing with the web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type: %T\n", response)

	defer response.Body.Close() // callers reponsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}
