package util

import (
	"file-manager/config"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type FileModel struct {
	Path       string
	Name       string
	Size       int64
	SizeStr    string
	Mode       os.FileMode
	ModTime    time.Time
	ModTimeStr string
	IsDir      bool
	IsImg      bool
}

// 读取文件夹
func ReadDir(dir string) ([]*FileModel, error) {
	fullPath := filepath.Join(config.RootPath, dir)
	fileInfos, err := ioutil.ReadDir(fullPath)
	if err != nil {
		panic("文件夹错误")
	}
	var fileSlice []*FileModel
	for _, fileInfo := range fileInfos {
		isImg, _ := regexp.MatchString("\\.(jpg|jpeg|bmp|png)$", fileInfo.Name())
		file := &FileModel{
			filepath.Join(dir, fileInfo.Name()),
			fileInfo.Name(),
			fileInfo.Size(),
			FormatFileSize(fileInfo.Size()),
			fileInfo.Mode(),
			fileInfo.ModTime(),
			FormatTime(fileInfo.ModTime()),
			fileInfo.IsDir(),
			isImg,
		}
		fileSlice = append(fileSlice, file)
	}
	return fileSlice, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func FormatFileSize(b int64) string {
	fb := float64(b)
	if fb < 1024 {
		return fmt.Sprintf("%0.2f B", fb)
	}
	kb := fb / 1024
	if kb < 1024 {
		return fmt.Sprintf("%0.2f KB", kb)
	}
	mb := kb / 1024
	if mb < 1024 {
		return fmt.Sprintf("%0.2f MB", mb)
	}
	gb := mb / 1024
	return fmt.Sprintf("%0.2f GB", gb)
}

func FormatTime(t time.Time) string {
	// year, month, day := t.Date()
	return t.Format("2006/01/01 15:04:05")
}
