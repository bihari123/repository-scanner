package main

import (
	"fmt"
	"log"

	gitutils "github.com/bihari123/repository-scanner/git-utils"
	readfile "github.com/bihari123/repository-scanner/read-file"
)

func main() {
	fmt.Println("\t\t\tWelcome to the repository scanne!\n\t\t\tThis is a tool that helps you secure your keys\nPlease make sure that your git configuration is set up on your system and you have the access to the repo")

	folder:=gitutils.GitClone("https://github.com/bihari123/My-Book-On-Docker-And-Kubernetes")
fmt.Println(folder)

	content, err := readfile.ReadFile("LICENSE")
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range content {
		fmt.Println(val)
	}
}
