package services

import (
	"os"
	"practical/internal/utils"
)

func ReadDirectory(dirPath string) ([]string, error) {
	if dirPath == "" {
		return nil, utils.NewCustomError(utils.ErrEmptyDirPath)
	}

	fileInfo, err := os.ReadDir(dirPath)

	if err != nil {
		return nil, utils.NewCustomError(utils.ErrReadingDir + err.Error())
	}

	var files []string
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	return files, nil
}
func ReadFileLines(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {

		return "", err
	}
	return string(content), nil

}
