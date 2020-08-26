package util

import (
	"file-manager/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileType uint32

const (
	FileTypeUnknown FileType = 0
	FileTypeImage FileType = 1
	FileTypeVideo FileType = 2
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
	IsVideo    bool
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
		isImg := false
		isVideo := false
		if !fileInfo.IsDir() {
			fileType := CheckFileType(filepath.Join(fullPath, fileInfo.Name()))
			if fileType == FileTypeImage {
				isImg = true;
			} else if (fileType == FileTypeVideo) {
				isVideo = true
			}
		}
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
			isVideo,
		}
		fileSlice = append(fileSlice, file)
	}
	return fileSlice, nil
}

func CheckFileType(filePath string) FileType {
	f, err := os.Open(filePath)
	if err != nil {
		return FileTypeUnknown
	}
	defer f.Close()
	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return FileTypeUnknown
	}
	contentType := http.DetectContentType(buffer)
	if strings.Contains(contentType, "image") {
		return FileTypeImage
	}
	if strings.Contains(contentType, "video") {
		return FileTypeVideo
	}
	return FileTypeUnknown
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
