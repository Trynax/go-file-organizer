package main


import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileType string

const (
	TypeImage FileType = "Image"
	TypeVideo FileType = "Video"
	TypeDocument FileType = "Document"
	TypeAudio FileType = "Audio"
	TypeArchive FileType = "Archive"
	TypeCode FileType = "Code"
	TypeOther FileType = "Other"

)

type FolderType string

const (
	TypeImageFolder FolderType = "Images"
	TypeVideoFolder FolderType = "Videos"
	TypeDocumentFolder FolderType = "Documents"
	TypeAudioFolder FolderType = "Audios"
	TypeArchiveFolder FolderType = "Archives"
	TypeCodeFolder FolderType = "Codes"
	TypeOtherFolder FolderType = "Others"

)

func GetFileType(filename string) FileType{

	ext := strings.ToLower(filepath.Ext(filename))
	switch ext{
		case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp":
			return TypeImage
		case ".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm":
			return TypeVideo
		case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt", ".csv" , ".odt", ".rtf":
			return TypeDocument
		case ".mp3", ".wav", ".flac", ".aac", ".ogg", ".wma":
			return TypeAudio
		case ".zip", ".rar", ".tar", ".gz", ".7z", ".bz2":
			return TypeArchive
		case ".c", ".cpp", ".h", ".java", ".py", ".js", ".html", ".css", ".go", ".rb":
			return TypeCode
		default:
			return TypeOther
}
}

func GetCategoryFolder(fileType FileType) FolderType{
	switch fileType{
		case TypeImage:
			return TypeImageFolder
		case TypeVideo:
			return TypeVideoFolder
		case TypeDocument:
			return TypeDocumentFolder
		case TypeAudio:
			return TypeAudioFolder
		case TypeArchive:
			return TypeArchiveFolder
		case TypeCode:
			return TypeCodeFolder
		default:
			return TypeOtherFolder
	}
}
func CreateParentFolder(sourceFolder string) (string, error){

	sourceInfo, err := os.Stat(sourceFolder)

	if err != nil{
		return "", fmt.Errorf("Error getting source folder info: %v", err)
	}

	if !sourceInfo.IsDir(){
		return "", fmt.Errorf("%s is not a directory", sourceFolder)
	}

	sourceAbspath, err := filepath.Abs(sourceFolder)

	if err!=nil{
		return "", fmt.Errorf("Error getting absolute path of source folder: %v", err)

	}

	parrentDir := filepath.Dir(sourceAbspath)
	organizedFolder := filepath.Join(parrentDir, "Organized-Files")


	err = os.MkdirAll(organizedFolder, 0755)

	if err != nil{
		return "", fmt.Errorf("Error creating organized folder: %v", err)
	}

	fmt.Println("Organized folder created at:", organizedFolder)
	return organizedFolder, nil

} 


func createCategoryFolder(organizedFolder string, category FolderType) (string, error){
	categoryFolder := filepath.Join(organizedFolder, string(category))

	err := os.MkdirAll(categoryFolder, 0755)
	
	if err != nil{
		return "", fmt.Errorf("Error creating category folder: %v", err)
	}

	return categoryFolder, nil

}


func HasFiles(folderPath string) (bool, error){
	files, err := os.ReadDir(folderPath)
	if err != nil{
		return false, err
	}

	if len(files) == 0 {
		return false, nil
	}
	return true, nil
}