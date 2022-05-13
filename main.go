package main

import (
	"fmt"
	"log"

	readfile "github.com/bihari123/repository-scanner/read-file"
)

func main() {
	content, err := readfile.ReadFile("LICENSE")
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range content {
		fmt.Println(val)
	}
}
