package controllers

import (
	"changeme/apps/ctx"
	"changeme/apps/datas"
	"changeme/apps/mysql"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"path/filepath"
)

// Files @Bean
type Files struct{}

func (f *Files) file() (string, error) {
	return ctx.OpenFileDialog(runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "",
		Title:                      "选中 binlog 原文件",
		Filters:                    nil,
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
}

func (f *Files) GetSystemFile() string {
	path, err := f.file()
	if err != nil {
		ctx.LogError("file err = " + err.Error())
		err.Error()
	}
	if path == "" {
		return ""
	}

	ctx.LogDebug("上传文件 path = " + path)
	err = mysql.SaveToSqlite(path)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (f *Files) GetDecodeRowCount() uint64 {
	return mysql.DecodeRowCount
}

func (f *Files) GetTitleList() []string {
	path := datas.GetSqlitePath("")
	arr, err := getDBFiles(path)
	if err != nil {
		return nil
	}
	got := make([]string, 0)
	for _, s := range arr {
		got = append(got, s)
	}
	return got
}

func getDBFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir) // 读取目录下所有文件
	if err != nil {
		return nil, err
	}

	var dbFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".db" {
			dbFiles = append(dbFiles, file.Name()) // 输出.db后缀的文件名
		}
	}

	return dbFiles, nil
}
