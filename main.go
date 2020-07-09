package main

import (
	commonController "crazyball/go-common/controller"
	"crazyball/go-common/httpServer"
	"file-manager/config"
	"file-manager/controller"
	"file-manager/util"
	"fmt"
	"html/template"
	"os"
)

func main() {
	fileInfo, err := os.Stat(config.RootPath)
	if err != nil || !fileInfo.IsDir() {
		fmt.Println("[ERROR] 文件夹路径不存在")
		panic("文件夹路径不存在")
	}
	// production config
	httpServer.Mode = httpServer.HttpServerModeProduction
	httpServer.LoggerFile = "/root/crazyball/static/logs/file-manager.log"
	httpServer.ErrorFile = "/root/crazyball/static/logs/file-manager-error.log"

	server := httpServer.New()
	server.HtmlDir = "view"
	server.StaticDir = "static"

	server.SetFuncMap(template.FuncMap{
		"formatFileSize": util.FormatFileSize,
		"formatTime":     util.FormatTime,
	})

	server.GET("/verify", commonController.VerifyTicket)
	server.POST("/api/upload",controller.UploadFile)
	server.POST("/upload", httpServer.ApiAuthMiddleware(""), controller.Upload)
	server.POST("/mkdir", httpServer.ApiAuthMiddleware(""), controller.CreateDir)
	server.POST("/remove", httpServer.ApiAuthMiddleware(""), controller.Delete)

	server.Use(httpServer.PageAuthMiddleware(""),controller.Index)

	server.Run(config.ServerPort)
}
