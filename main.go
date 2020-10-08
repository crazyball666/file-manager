package main

import (
	"crazyball/go-common/CBServer"
	"file-manager/config"
	"file-manager/router"
	"fmt"
	"os"
)

func init() {
	CBServer.Mode = CBServer.ModeProduction
}

func main() {
	fileInfo, err := os.Stat(config.RootPath)
	if err != nil || !fileInfo.IsDir() {
		fmt.Println("[ERROR] 文件夹路径不存在")
		panic("文件夹路径不存在")
	}

	server := CBServer.New();
	router.UseRoute(server)
	server.RunOnPort(config.ServerPort);
}
