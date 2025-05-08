package main


import(
	"fmt"
	"os"
	"path/filepath"
	"flag"
)


func main(){

	sourcePath := flag.String("path", ".", "Path to organize files from")
	recursive := flag.Bool("recursive", false, "Process subdirctories recursively")
	dryRun := flag.Bool("dry-run", false, "show what would be done without making chances")

	flag.Parse()


	absPath, err := filepath.Abs(*sourcePath)
	if err != nil{
		fmt.Println("Error getting absolute path:", err)
		os.Exit(1)

	}

	fmt.Printf("Starting file organization in: %s\n", absPath)
	fmt.Printf("Recursive mode: %v\n", *recursive)
	fmt.Printf("Dry run mode: %v\n", *dryRun)


	fmt.Println("File organization completed!")





}