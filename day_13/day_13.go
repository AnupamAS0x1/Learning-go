package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&payment"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)

	fmt.Println(result.Path)

	fmt.Println(result.Port())

	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("type of qurey params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}
	partsofurl := &url.URL{ // remember to pass &
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user+hitesh",
	}
	anotherURL := partsofurl.String()
	fmt.Println(anotherURL)

}
