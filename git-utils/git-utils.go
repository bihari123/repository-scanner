package gitutils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func GitClone(url string) (folderName string) {

	home, _ := os.UserHomeDir()
	//change the working directory of the program
	// err := os.Chdir(filepath.Join(home, "Desktop"))
	//  if err != nil {
	//  	panic(err)
	//  }
	path_to_desktop := filepath.Join(home, "Desktop")
	folderName = strings.SplitAfter(url, "/")[4]
		

	if exists(filepath.Join(path_to_desktop,folderName)) {
      log.Print("\n\ndirectory already exist, deleting........\n\n")
			cmd1 := exec.Command("rm" ,"-r", "-f", filepath.Join(path_to_desktop,folderName))
			err:=cmd1.Run()
			if err!=nil{
				log.Println(fmt.Errorf("error while deletting the folder %w",err))
			}
    time.Sleep(10000) 
		} 
 


 		cmd := exec.Command("git", "clone", url)
	cmd.Dir = path_to_desktop

	err := cmd.Run()
	os.Chdir(path_to_desktop)
	if err != nil {

				log.Fatal("Error in cloning repo", err, "Either try deleting the repo yourself(if already cloned )or run this on linux debian based distro")
		

	}

	err = os.Chdir(filepath.Join(path_to_desktop, folderName))
	if err != nil {
		log.Fatal(" can't go to the project directory after download", err)

	}

	return
}


// exists returns whether the given file or directory exists
func exists(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { 
    	return false 
    }else{
    	log.Fatal("some error happened at checking the existence of the directory ",err)
    }
    
    return false
}
