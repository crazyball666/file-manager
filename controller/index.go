package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"file-manger-go/util"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	path := filepath.Join(util.RootPath, c.Request.URL.Path)
	fmt.Println(path)
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		c.String(404, "404 not Found")
	} else {
		res, _ := util.ReadDir(path)
		// c.JSON(200, res)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"dirList": res,
		})
	}
}

// 上传多文件
func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	uploadPath := c.PostForm("uploadPath")
	if uploadPath == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "缺少上传Path",
		})
		return
	}
	files := form.File["files"]
	fmt.Println(uploadPath, files)
	if len(files) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "缺少上传文件",
		})
		return
	}
	for _, file := range files {
		filePath := path.Join(util.RootPath, uploadPath, file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": fmt.Sprintf("上传文件%s错误", file.Filename),
			})
			fmt.Println("【ERROR】", err)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "上传成功",
	})
}

// 创建文件夹
func CreateDir(c *gin.Context) {
	defer func() {
		c.Request.Body.Close()
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": err,
			})
		}
	}()
	var postJson map[string]interface{}
	postData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(postData, &postJson)
	if err != nil {
		panic("参数错误")
	}
	if _, ok := postJson["path"].(string); !ok {
		panic("缺少Path")
	}
	dirPath := path.Join(util.RootPath, postJson["path"].(string))
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		panic("创建文件夹失败: " + err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

// 删除文件
func Delete(c *gin.Context) {
	defer func() {
		c.Request.Body.Close()
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": err,
			})
		}
	}()
	var postJson map[string]interface{}
	postData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(postData, &postJson)
	if err != nil {
		panic("参数错误")
	}
	removePath, ok := postJson["path"].(string)
	if !ok {
		panic("缺少Path")
	}
	fullPath := path.Join(util.RootPath, removePath)
	if err := os.Remove(fullPath); err != nil {
		panic("删除失败: " + err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

func FormateSize(b int64) string {
	if b < 1024 {
		return fmt.Sprintln("%d B", b)
	}
	kb := b / 1024
	if kb < 1024 {
		return fmt.Sprintln("%f KB", kb)
	}
	mb := kb / 1024
	if mb < 1024 {
		return fmt.Sprintln("%f MB", mb)
	}
	gb := mb / 1024
	return fmt.Sprintln("%f GB", gb)
}
