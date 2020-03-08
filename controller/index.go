package controller

import (
	"file-manager/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	path := filepath.Join(util.RootPath, c.Request.URL.Path)
	fmt.Println(path)
	info, err := os.Stat(path)
	if err != nil {
		c.String(404,"404 not found");
	}else if !info.IsDir() {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err.Error())
		}
		c.Writer.Header().Set("Content-Type","application/octet-stream")
		c.Status(200)
		c.Writer.Write(buf)
	} else {
		res, _ := util.ReadDir(c.Request.URL.Path)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"basePath":c.Request.URL.Path,
			"dirList": res,
		})
	}
}

// 上传多文件
func Upload(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
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
		filePath := path.Join(util.RootPath, basePath, file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			panic(fmt.Sprintf("上传文件%s错误", file.Filename))
		}
	}
	c.Redirect(302,basePath)
}


// 创建文件夹
func CreateDir(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": err,
			})
		}
	}()
	basePath := c.PostForm("basePath")
	dirName:= c.PostForm("dirName")
	if len(basePath) == 0 || len(dirName) == 0 {
		panic("不能为空")
	}
	dirPath := path.Join(util.RootPath, basePath,dirName)
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		panic("创建文件夹失败: " + err.Error())
	}
	c.Redirect(302,basePath)
}

// 删除文件
func Delete(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": err,
			})
		}
	}()
	basePath := c.PostForm("basePath")
	fileName:= c.PostForm("fileName")
	if len(basePath) == 0 || len(fileName) == 0 {
		panic("不能为空")
	}
	fullPath := path.Join(util.RootPath, basePath,fileName)
	if err := os.Remove(fullPath); err != nil {
		panic("删除失败: " + err.Error())
	}
	c.Redirect(302,basePath)
}

