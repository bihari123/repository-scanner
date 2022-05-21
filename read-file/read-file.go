package readfile

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) (content []string, err error) {
	/*The simplest way of reading a text
	  or binary file in Go is to use the ReadFile() function
	  from the os package. This function reads the entire content
	  of the file into a byte slice, so you should be careful when
	  trying to read a large file - in this case,
	  you should read the file line by line or in chunks.
	  For small files, this function is more than sufficient.
	*/
	f, err := os.Open(fileName)
	// remember to close the file at the end
	defer f.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return

}
func ReadRepo(fileTree []string)(warnings []string){
	for _,filePath:= range fileTree{
		dir,err:= os.Stat(filePath)
	
		if dir.IsDir() || strings.Contains(filePath,".git") {
         continue 
		}
    fileContent,err:=ReadFile(filePath)
    if err!=nil{
    	log.Println("can't read the file",filePath)
    }

    for _,word:= range fileContent{
         if strings.Contains(word,"public_key") || strings.Contains(word, "private_key"){
             warnings=append(warnings,"SECRET_KEY found at ",filePath)
         }  
    }
	}

	return
}
