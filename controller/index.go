package controller

import (
	"crazyball/go-common/httpServer"
	"file-manager/config"
	"file-manager/util"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func Index(c *httpServer.HttpContext) {
	fullpath := filepath.Join(config.RootPath, c.Request.URL.Path)
	info, err := os.Stat(fullpath)
	if err != nil {
		c.String(404, "404 not found");
	} else if !info.IsDir() {
		file, err := os.Open(fullpath)
		if err != nil {
			c.Error("open file is error");
		}
		http.ServeContent(c.Writer, c.Request, fullpath, info.ModTime(), file)
	} else {
		res, _ := util.ReadDir(c.Request.URL.Path)
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"basePath": c.Request.URL.Path,
			"dirList":  res,
		})
	}
}

// 上传多文件【通用接口】
func UploadFile(c *httpServer.HttpContext) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    500,
				"message": err,
			})
		}
	}()
	form, _ := c.MultipartForm()
	files := form.File["files"]
	if len(files) < 1 {
		panic("缺少上传文件")
	}
	data := make([]string, 0)

	now := time.Now()
	dateStr := now.Format("2006-01-02")

	uploadDir := path.Join(config.StaticPath, config.UploadDirName, dateStr)
	if !util.Exists(uploadDir) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			panic("创建文件夹失败")
		}
	}
	for _, file := range files {
		fileName := fmt.Sprintf("%d_", now.Unix()) + file.Filename
		filePath := path.Join(uploadDir, fileName)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			panic(fmt.Sprintf("上传文件%s错误", file.Filename))
		} else {
			data = append(data, path.Join(config.StaticHost, config.UploadDirName, dateStr, fileName))
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func Upload(c *httpServer.HttpContext) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    500,
				"message": err,
			})
		}
	}()
	form, _ := c.MultipartForm()
	basePath := c.PostForm("basePath")
	if basePath == "" {
		panic("缺少上传Path")
	}
	files := form.File["files"]
	if len(files) < 1 {
		panic("缺少上传文件")
	}
	for _, file := range files {
		filePath := path.Join(config.RootPath, basePath, file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			panic(fmt.Sprintf("上传文件%s错误", file.Filename))
		}
	}
	c.Success(nil)
}

// 创建文件夹
func CreateDir(c *httpServer.HttpContext) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    500,
				"message": err,
			})
		}
	}()
	basePath := c.PostForm("basePath")
	dirName := c.PostForm("dirName")
	if len(basePath) == 0 || len(dirName) == 0 {
		panic("不能为空")
	}
	dirPath := path.Join(config.RootPath, basePath, dirName)
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		panic("创建文件夹失败: " + err.Error())
	}
	c.Redirect(302, basePath)
}

// 删除文件
func Delete(c *httpServer.HttpContext) {
	filePath := c.Query("path")
	fmt.Println(c.Query("path"))
	if len(filePath) == 0 {
		c.Error("路径不能为空")
		return
	}
	fullPath := path.Join(config.RootPath, filePath)
	if err := os.Remove(fullPath); err != nil {
		c.Error("删除失败: " + err.Error())
		return
	}
	c.Success(nil);
}

// 查看文件内容接口
func GetFileContent(c *httpServer.HttpContext) {
	relativePath := c.Query("path")
	fullPath := filepath.Join(config.RootPath, relativePath)
	info, err := os.Stat(fullPath)
	if err != nil {
		c.Error("file is not exist");
	} else if info.IsDir() {
		c.Error("file is dir");
	} else {
		fileType := util.CheckFileType(fullPath)
		if info.Size() > 5*1024*1024 && fileType == util.FileTypeUnknown {
			c.Error("file size is larger than 5 mb")
			return
		}
		file, err := os.Open(fullPath)
		if err != nil {
			c.Error("open file is error");
			return
		}
		http.ServeContent(c.Writer, c.Request, fullPath, info.ModTime(), file)
	}
}
