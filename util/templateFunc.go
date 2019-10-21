package util

import (
	"fmt"
	"time"
)

/**
模板函数
*/

func FormatFileSize(b int64) string {
	fb := float64(b)
	if fb < 1024 {
		return fmt.Sprintf("%0.2f B", fb)
	}
	kb := fb / 1024
	if kb < 1024 {
		return fmt.Sprintf("%0.2f KB", kb)
	}
	mb := kb / 1024
	if mb < 1024 {
		return fmt.Sprintf("%0.2f MB", mb)
	}
	gb := mb / 1024
	return fmt.Sprintf("%0.2f GB", gb)
}

func FormatTime(t time.Time) string {
	// year, month, day := t.Date()
	return t.Format("2006/01/01 15:04:05")
}
