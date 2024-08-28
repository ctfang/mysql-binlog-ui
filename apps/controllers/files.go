package controllers

import (
	"changeme/apps/ctx"
	"changeme/apps/datas"
	"changeme/apps/mysql"
	"changeme/apps/orm"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
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
	return path
}

func (f *Files) SaveToSqlite(path string) string {
	err := mysql.SaveToSqlite(path)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (f *Files) GetDecodeRowCount() uint64 {
	return mysql.DecodeRowCount
}

func (f *Files) ClearAllData() {
	path, _ := datas.GetDataDir("mysql-binlog-ui/")
	path = datas.ToPath(path)

	// 删除所有目录
	os.RemoveAll(path)
}

type TitleData struct {
	List  []*orm.UploadLogs
	Total int64
}

// GetTitleList 获取标题列表
func (f *Files) GetTitleList(page int, limit int) TitleData {
	got, count := orm.NewOrmUploadLogs().Paginate(page, limit)

	for _, logs := range got {
		logs.FileSize = float64(int(logs.FileSize))
	}

	return TitleData{
		List:  got,
		Total: count,
	}
}

func (f *Files) DeleteFile(ID int64) {
	file, ok := orm.NewOrmUploadLogs().WhereId(ID).First()
	if ok {
		orm.NewOrmUploadLogs().WhereId(ID).Delete()

		err := os.Remove(file.Database)
		if err != nil {
			ctx.LogDebug("del err = ", err)
		}
	}
}
