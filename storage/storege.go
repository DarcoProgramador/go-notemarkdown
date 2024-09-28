package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

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
