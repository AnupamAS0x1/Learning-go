package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("playing with POST/GET request")
	PerformPostRequest()
}

func PerformGedRequest() {

	const myurl = "https://github.com/Snapchat"
	response, err := http.Get(myurl)

	if err != nil {

		panic(err)

	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("content length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostRequest() {

	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`

		{
			"coursename":"lets go with golang",
			"price": 0,
			"platform":"learncodeonline.in"	
		}

	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {

		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
