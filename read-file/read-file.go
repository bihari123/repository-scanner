package readfile

import (
	"bufio"
	"log"
	"os"
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
