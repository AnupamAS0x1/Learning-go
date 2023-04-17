package main

import (
	"encoding/json"
	"fmt"
)

type course struct { //playing with aliases
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // the - here means dont through this field in response
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("creating json data ")
	EncodeJson()
}

func EncodeJson() {

	lcocources := []course{

		{"ReactJS bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN bootcamp", 199, "learncodeonline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular bootcamp", 299, "le arncodeonline.in", "hit123", nil},
	}
	// pacakge this data as JSON data
	finaljson, err := json.MarshalIndent(lcocources, "", "\t")

	if err != nil {
		panic(err)

	}
	fmt.Printf("%s\n", finaljson)
}
