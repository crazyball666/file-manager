package schedule

import (
	"crazyball/go-common/Crontab"
	"file-manager/config"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func init() {
	id, err := Crontab.Manager().AddFunc("0 0 */6 * * *", clearTempDir)
	fmt.Println(id, err)
}

func clearTempDir() {
	dirs, err := ioutil.ReadDir(path.Join(config.RootPath, config.TempDirName))
	if err != nil {
		return
	}
	for _, dir := range dirs {
		_ = os.RemoveAll(path.Join(config.RootPath, config.TempDirName, dir.Name()))
	}
	fmt.Println(time.Now(), "清空temp目录")
}
