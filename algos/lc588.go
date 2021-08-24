package algos

import (
	"sort"
	"strings"
)

type FileSystem struct {
	File
}

const (
	FILE int = iota
	FOLDER
)

type File struct {
	Name     string
	Type     int
	Data     string
	Subfiles map[string]*File
}

func Constructor() FileSystem {
	return FileSystem{
		File{
			Name:     "/",
			Type:     FOLDER,
			Subfiles: make(map[string]*File),
		},
	}
}

func (this *FileSystem) Ls(path string) []string {
	finalDir := &this.File
	files := strings.Split(path, "/")
	for _, file := range files {
		if file == "" {
			continue
		}
		finalDir = finalDir.Subfiles[file]
	}
	if finalDir.Type == FILE {
		return []string{finalDir.Name}
	}
	result := make([]string, 0, len(finalDir.Subfiles))
	for _, subfile := range finalDir.Subfiles {
		result = append(result, subfile.Name)
	}
	sort.Strings(result)
	return result
}

func (this *FileSystem) Mkdir(path string) {
	folderNames := strings.Split(path, "/")
	finalDir := &this.File
	for _, folderName := range folderNames {
		if folderName == "" {
			continue
		}
		_, ok := finalDir.Subfiles[folderName]
		if !ok {
			finalDir.Subfiles[folderName] = &File{
				Name:     folderName,
				Type:     FOLDER,
				Subfiles: make(map[string]*File),
			}
		}
		finalDir = finalDir.Subfiles[folderName]
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	fileNames := strings.Split(filePath, "/")
	finalFile := &this.File
	for i := 0; i < len(fileNames)-1; i++ {
		fileName := fileNames[i]
		if fileName == "" {
			continue
		}
		finalFile = finalFile.Subfiles[fileName]
	}
	file, ok := finalFile.Subfiles[fileNames[len(fileNames)-1]]
	if ok {
		file.Data += content
	} else {
		finalFile.Subfiles[fileNames[len(fileNames)-1]] = &File{
			Name: fileNames[len(fileNames)-1],
			Data: content,
			Type: FILE,
		}
	}
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	fileNames := strings.Split(filePath, "/")
	finalFile := &this.File
	for _, fileName := range fileNames {
		if fileName == "" {
			continue
		}
		finalFile = finalFile.Subfiles[fileName]
	}
	return finalFile.Data
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
