package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("anu")
	content := "this needs to go in a file - learncodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(nil)

	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(nil)

	}

	fmt.Println("length is :", length)
	defer file.Close()
	readfile("./mylcogofile.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("text data inside th file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {

		panic(err)
	}
}
