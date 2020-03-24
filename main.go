package main

import (
	"file-manager/config"
	"file-manager/controller"
	"file-manager/util"
	"fmt"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fileInfo, err := os.Stat(config.RootPath)
	if err != nil || !fileInfo.IsDir() {
		fmt.Println("[ERROR] 文件夹路径不存在")
		panic("文件夹路径不存在")
	}

	app := gin.Default()
	app.Static("/static", "./static")

	app.SetFuncMap(template.FuncMap{
		"formatFileSize": util.FormatFileSize,
		"formatTime":     util.FormatTime,
	})
	app.LoadHTMLGlob("view/*")

	app.POST("/api/upload",controller.UploadFile)

	app.POST("/upload", controller.Upload)
	app.POST("/mkdir", controller.CreateDir)
	app.POST("/remove", controller.Delete)

	app.Use(controller.Index)

	app.Run(config.ServerPort)
}
