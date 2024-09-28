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
