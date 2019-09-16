package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path, _ := os.Getwd()
	filePath := path + "/Basics/readFile/text.txt"
	data, err := ioutil.ReadFile(filePath)
	checkError(err)
	fmt.Println("File Content: " + string(data))

	//Read file in byte chunk
	f, err := os.Open(filePath)
	checkError(err)
	b1 := make([]byte, 10)
	n1, err := f.Read(b1)
	checkError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	f.Close()
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
