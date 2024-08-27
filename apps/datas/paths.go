package datas

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func ToPath(savePath string) string {
	if runtime.GOOS != "windows" {
		return savePath
	}
	savePath = strings.ReplaceAll(savePath, "\\", "/")
	savePath = savePath[:3] + strings.ReplaceAll(savePath[3:], ":", "/")
	savePath = strings.ReplaceAll(savePath, "//", "/")
	savePath = strings.ReplaceAll(savePath, "-", "_")

	return savePath
}

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
func GetPath(filePath string) string {
	dir, _ := GetDataDir("mysql-binlog-ui/" + filePath)
	dir = strings.ReplaceAll(dir, "//", "/")
	savePath := path.Dir(dir)
	if savePath == "." {
		dir = strings.ReplaceAll(dir, "\\", "/")
	}
	return dir
}

// GetSqlitePath 根据系统类型, 返回对应保存数据的目录
func GetSqlitePath(path string) string {
	return GetPath("/sqlite/" + path)
}

// GetFileSizeMB 返回文件大小，单位为MB
func GetFileSizeMB(filePath string) float64 {
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0
	}

	// 获取文件大小，并转换为MB
	fileSizeMB := float64(fileInfo.Size()) / 1024 / 1024
	return fileSizeMB
}
