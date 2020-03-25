package main

import (
	commonController "crazyball/go-common/controller"
	commonMiddleware "crazyball/go-common/middleware"
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

	app.Use(commonMiddleware.Cors)

	app.Static("/static", "./static")

	app.SetFuncMap(template.FuncMap{
		"formatFileSize": util.FormatFileSize,
		"formatTime":     util.FormatTime,
	})
	app.LoadHTMLGlob("view/*")


	app.GET("/verify", commonController.VerifyTicket)

	app.POST("/api/upload",controller.UploadFile)

	app.POST("/upload", commonMiddleware.AuthPage(""), controller.Upload)
	app.POST("/mkdir", commonMiddleware.AuthPage(""), controller.CreateDir)
	app.POST("/remove", commonMiddleware.AuthPage(""), controller.Delete)

	app.Use(commonMiddleware.AuthPage(""),controller.Index)

	app.Run(config.ServerPort)
}
