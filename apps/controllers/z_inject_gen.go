// gen for home toolset
package controllers

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _BinlogSingle *Binlog
var _FilesSingle *Files
var _HelpSingle *Help

func GetAllProvider() []interface{} {
	return []interface{}{
		NewBinlog(),
		NewFiles(),
		NewHelp(),
	}
}

func NewBinlog() *Binlog {
	if _BinlogSingle == nil {
		_BinlogSingle = &Binlog{}
		providers.AfterProvider(_BinlogSingle, "")
	}
	return _BinlogSingle
}
func NewFiles() *Files {
	if _FilesSingle == nil {
		_FilesSingle = &Files{}
		providers.AfterProvider(_FilesSingle, "")
	}
	return _FilesSingle
}
func NewHelp() *Help {
	if _HelpSingle == nil {
		_HelpSingle = &Help{}
		providers.AfterProvider(_HelpSingle, "")
	}
	return _HelpSingle
}
