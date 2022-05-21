package main

import (
	"fmt"
	"log"

	gitutils "github.com/bihari123/repository-scanner/git-utils"
	readfile "github.com/bihari123/repository-scanner/read-file"
)

func main() {
	fmt.Println("\t\t\tWelcome to the repository scanne!\n\t\t\tThis is a tool that helps you secure your keys\nPlease make sure that your git configuration is set up on your system and you have the access to the repo")
	//
	log.Println("cloning repository to ~/Desktop.....")
	folder := gitutils.GitClone("https://github.com/bihari123/My-Book-On-Docker-And-Kubernetes")
	log.Println("git repository cloned at ~/Desktop/", folder)

	filePaths := readfile.WalkThroughoutRepo()

	log.Println("fetched the following files\n", filePaths)

	log.Println(" Scanning the files.....")

	warnings := readfile.ReadRepo(filePaths)
	fmt.Println("#######################################################")
	fmt.Println("Scan Result")

	if len(warnings) > 0 {
		fmt.Println(warnings)
	} else {
		fmt.Println("No SECRET_KEY found")
	}

	fmt.Println("############################################################")

}
