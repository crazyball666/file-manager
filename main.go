package main

import (
	"file-manager/controller"
	"file-manager/util"
	"fmt"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Printf("输入监听端口号： ")
	// fmt.Scanln(&util.ServerPort)
	// fmt.Printf("输入文件夹路径： ")
	// fmt.Scanln(&util.RootPath)

	fileInfo, err := os.Stat(util.RootPath)
	if err != nil || !fileInfo.IsDir() {
		fmt.Println("[ERROR] 文件夹路径不存在,enter键退出")
		fmt.Scanln()
		panic("文件夹路径不存在")
	}

	app := gin.Default()
	app.Static("/static", "./static")

	app.SetFuncMap(template.FuncMap{
		"formatFileSize": util.FormatFileSize,
		"formatTime":     util.FormatTime,
	})
	app.LoadHTMLGlob("view/*")

	app.POST("/upload", controller.Upload)
	app.POST("/mkdir", controller.CreateDir)
	app.POST("/remove", controller.Delete)
	app.Use(controller.Index)
	app.Run(util.ServerPort)
}
