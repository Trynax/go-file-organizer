package main


import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileType string;

const (
	TypeImage FileType = "Images"
	TypeVideo FileType = "Videos"
	TypeDocument FileType = "Documents"
	TypeAudio FileType = "Audio"
	TypeArchive FileType = "Archives"
	TypeCode FileType = "Code"
	TypeOther FileType = "Others"

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