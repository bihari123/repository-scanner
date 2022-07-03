package main

import (
	"fmt"

	readfile "github.com/bihari123/repository-scanner/read-file"
	"github.com/bihari123/repository-scanner/structs"
)

func main() {

	input:=structs.GitHub{
		Owner: "bihari123",
		Repository: "My-Book-On-Docker-And-Kubernetes",

	}

	content,_:=readfile.GetDirContent(input.Owner,input.Repository)

	fmt.Println(content)
}
