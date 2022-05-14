package gitutils

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GitClone(url string) (folderName string) {

	home, _ := os.UserHomeDir()
	// change the working directory of the program
	//err := os.Chdir(filepath.Join(home, "Desktop"))
	// if err != nil {
	// 	panic(err)
	// }
	cmd := exec.Command("git", "clone", url)
	cmd.Dir = filepath.Join(home, "Desktop")
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	return url
}
