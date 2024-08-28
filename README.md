## About

读取原始的 MySQL Binlog文件后，生成更加可读的结构化文件。 解析内容后，可以在 UI 上直接便捷的查看。
也可以通过生成的 sqlite 数据库文件，直接在数据库客户端上查看。

### Prerequisites

* Go (latest version)
* Node.js >= 16
* NPM >= 9

### Compile and run
[CONTRIBUTING.md](..%2F..%2Ftiny-rdm%2F.github%2FCONTRIBUTING.md)
```bash
wails dev
```

### 按照固定版本 Wails CLI v2.8.0
go install github.com/wailsapp/wails/v2/cmd/wails@v2.8.0


wails build  -platform  windows/amd64
wails build  -platform  darwin/universal
wails build  -platform  linux/amd64