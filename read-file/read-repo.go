package readfile

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	gitutils "github.com/bihari123/repository-scanner/git-utils"
)

func GetDirContent(owner, repository string) (warnings []string,err error) {
	gitClient := gitutils.GitAuthClient(os.Getenv("GITHUB_PERSONAL_TOKEN"))

	_, dir, _, err := gitClient.Repositories.GetContents(context.Background(), owner, repository, "/", nil)

	if err != nil {
		fmt.Printf("Error in getting the directory content: %v", err)
		return []string{},err 
	}

	for _, files := range dir {

		fileContent, _, _, err := gitClient.Repositories.GetContents(context.Background(), owner, repository, *files.Path, nil)

		if err != nil {
			fmt.Printf("Error getting fileContent of %s : %v", *files.Path, err)
			return []string{},err 
		}


		decodedContent,err:=base64.StdEncoding.DecodeString(*fileContent.Content)

		if err!=nil{
			fmt.Printf("error decoding the content of filepath: %s \n %v",*files.Path,err)
			return []string{},err 
		}
 for _,word:= range string(decodedContent){
           if strings.Contains(word,"public_key") || strings.Contains(word, "private_key"){
               warnings=append(warnings,"SECRET_KEY found at ",*files.Path)
           }

	}

fmt.Println(string(decodedContent))
}
	
	return 
}

