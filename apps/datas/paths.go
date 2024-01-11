package datas

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetDataDir(appName string) (string, error) {
	var dataDir string

	switch runtime.GOOS {
	case "windows":
		// 在 Windows 上使用 APPDATA 环境变量
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", os.ErrNotExist
		}
		dataDir = filepath.Join(appData, appName)

	case "darwin":
		// 在 macOS 上使用 HOME 环境变量并追加通用路径
		home := os.Getenv("HOME")
		if home == "" {
			return "", os.ErrNotExist
		}
		dataDir = filepath.Join(home, "Library", "Application Support", appName)

	default:
		// 对于其他操作系统，可以返回一个错误或提供另一种路径
		return "", os.ErrNotExist
	}

	return dataDir, nil
}

// GetPath 根据系统类型, 返回对应保存数据的目录
func GetPath(path string) string {
	dir, _ := GetDataDir("mysql-binlog-ui/" + path)
	return strings.ReplaceAll(dir, "//", "/")
}

// GetSqlitePath 根据系统类型, 返回对应保存数据的目录
func GetSqlitePath(path string) string {
	return GetPath("/sqlite/" + path)
}
