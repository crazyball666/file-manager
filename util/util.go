package util

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

// 读取文件夹
func ReadDir(dir string) ([]map[string]interface{}, error) {
	fullPath := filepath.Join(RootPath,dir)
	fileInfos, err := ioutil.ReadDir(fullPath)
	if err != nil {
		panic("文件夹错误")
	}
	var fileSlice []map[string]interface{}
	for _, fileInfo := range fileInfos {
		isImg, _ := regexp.MatchString("\\.(jpg|jpeg|bmp|png)$", fileInfo.Name())

		fileMap := map[string]interface{}{
			"path":        filepath.Join(dir,fileInfo.Name()),
			"name":        fileInfo.Name(),
			"size":        fileInfo.Size(),
			"mode":        fileInfo.Mode(),
			"modify_time": fileInfo.ModTime(),
			"is_dir":      fileInfo.IsDir(),
			"is_img":      isImg,
		}
		fileSlice = append(fileSlice, fileMap)
	}
	return fileSlice, nil
}
