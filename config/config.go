package config

/**
 根路径（自定义）
 RootPath
	- static 静态资源服务
		- uploads 持久化上传文件
	- temp   临时文件
	- logs   日志文件
*/

var (
	//RootPath   string = "/Users/crazyball/Desktop"
	RootPath      string = "/root/crazyball" // 根路径
	ServerPort    int    = 8080              //端口号
	TempDirName   string = "temp"            // 临时文件夹
	StaticDirName string = "static"          // 静态资源地址
	UploadDirName string = "uploads"         // 上传文件夹
	StaticHost    string = "static.crazyball.xyz"
)
