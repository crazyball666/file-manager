package router

import (
	"crazyball/go-common/CBServer"
	publicController "crazyball/go-common/CBServer/controller"
	publicMiddleware "crazyball/go-common/CBServer/middleware"
	"file-manager/controller"
)

func UseRoute(app *CBServer.Server) {
	/**
	 * 中间件
	 */
	app.Use(publicMiddleware.NewLoggerMiddleware("/root/crazyball/logs/file.log"))
	app.Use(publicMiddleware.NewRecoverMiddleware("/root/crazyball/logs/file-error.log"))
	app.Use(publicMiddleware.CorsMiddleware)

	app.Static("/_static", "./static")
	app.LoadHTMLGlob("./view/*")

	app.GET("/verify", publicController.VerifyTicket)

	// 通用接口
	app.POST("/api/upload", CBServer.WithCBContext(controller.UploadFile))
	app.POST("/api/temp/upload", CBServer.WithCBContext(controller.UploadTempFile))

	/// 文件系统操作接口
	app.GET(
		"/getFileDetail",
		publicMiddleware.PageAuthMiddleware(""),
		CBServer.WithCBContext(controller.GetFileContent),
	)
	app.POST(
		"/upload",
		publicMiddleware.PageAuthMiddleware(""),
		CBServer.WithCBContext(controller.Upload),
	)
	app.POST(
		"/mkdir",
		publicMiddleware.PageAuthMiddleware(""),
		CBServer.WithCBContext(controller.CreateDir), )
	app.GET(
		"/remove",
		publicMiddleware.PageAuthMiddleware(""),
		CBServer.WithCBContext(controller.Delete),
	)

	/// 文件列表页面
	app.Use(
		publicMiddleware.PageAuthMiddleware(""),
		CBServer.WithCBContext(controller.FileListPage),
	)
}
