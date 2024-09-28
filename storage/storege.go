package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

// SaveFile saves a file to the disk and returns the file path and an error
// file is a pointer to a multipart.FileHeader
// basePath is the path where the file will be saved
func SaveFile(file *multipart.FileHeader, basePath string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filePath := fmt.Sprintf("%s/%s", basePath, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}

// DeleteFile deletes a file from the disk and returns an error
// filePath is the path to the file
func DeleteFile(fileName, basePath string) error {
	return os.Remove(basePath + "/" + fileName)
}

// GetFiles returns a list of files in a directory
// basePath is the path to the directory
func GetFiles(basePath string) ([]string, error) {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

// GetFile returns a file from the disk
// filePath is the path to the file
func GetFile(filePath string) (*os.File, error) {
	return os.Open(filePath)
}
