package readfile

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func WalkThroughoutRepo()(filenames []string){

  err := filepath.Walk(".",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path, info.Size())
    filenames=append(filenames,path)
    return nil
})
if err != nil {
    log.Println(err)

	}

	return
}
