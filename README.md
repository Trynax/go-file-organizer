# File Organizer

A simple Go utility to organize files into categorized folders.

## Features

- Organizes files from a specified directory into category-based folders
- Preserves original files
- Easy to use with command-line flags
- Automatically creates an organized directory structure

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/Trynax/go-file-organizer.git
   cd go-file-organizer
   ```

2. Build the application:
   ```
   go build
   ```

## Usage

Run the application with the following command:

```
./file-organizer --path="/path/to/your/files"
```

### Options

- `--path`: Specifies the directory containing files to organize (default: current directory)

### Example

```
./file-organizer --path="/Users/username/Downloads"
```

This will:
1. Check if the specified directory has files
2. Create an "organized-Files" directory
3. Organize files by type into appropriate subdirectories
4. Display a completion message when done

## File Organization Categories

Files are organized into the following categories:
- Documents (PDF, DOC, DOCX, TXT, etc.)
- Images (JPG, PNG, GIF, etc.)
- Audio (MP3, WAV, etc.)
- Video (MP4, MOV, etc.)
- Archives (ZIP, RAR, etc.)
- Others (Any file type not covered by the above categories)

## Notes

- Original files remain untouched
- The program creates a new directory structure with organized files

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.