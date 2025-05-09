package main


import(
	"fmt"
	"os"
	"path/filepath"
	"flag"
)


func main(){

	sourcePath := flag.String("path", ".", "Path to organize files from")
	flag.Parse()
	absPath, err := filepath.Abs(*sourcePath)
	if err != nil{
		fmt.Println("Error getting absolute path:", err)
		os.Exit(1)
	}
	
	// check if the folder has files
	_, err = HasFiles(absPath)
	if err != nil{
		fmt.Println("Error checking for files:", err)
		os.Exit(1)
	}

	organizedFolder,err :=CreateOrganizedFolder(absPath)
	if err != nil{
		fmt.Println("Error creating organized folder:", err)
		os.Exit(1)
	}

	// create category folders


	
OragnizeFiles(absPath, organizedFolder)



	fmt.Printf("Starting file organization in: %s\n", absPath)



	fmt.Println("File organization completed!")





}

