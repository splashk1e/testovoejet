package fileservice

import (
	"io"
	"os"
	"strings"
)

type IFileService interface {
	ReadFile() (string, error)
	WriteFile(text string) error
}
type FileService struct {
	filePath string
}

func NewFileService(filePath string) *FileService {
	return &FileService{
		filePath: filePath,
	}
}
func (s *FileService) ReadFile() (string, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	data := make([]byte, 64)
	var builder strings.Builder
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		builder.Write(data[:n])
	}
	return builder.String(), nil
}
func (s *FileService) WriteFile(text string) error {
	file, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}
	if _, err := file.WriteString(text); err != nil {
		return err
	}
	return nil
}
