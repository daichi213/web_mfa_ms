# エラー

## Gin 関連

### import cycle not allowed

Gin を使用して、main.go を起動した際に以下のような循環エラーが発生した。

```bash
/go/src/api # go run main.go
package command-line-arguments
        imports api
        imports api: import cycle not allowed
```

原因は、module 直下の`package api`のさらに下に package を配置していることだった。

```bash
/go/src/api # tree -L 3 ./
./
├── Dockerfile
├── api
│   ├── constants.go
│   ├── controllers
│   │   ├── controller.go
│   │   ├── controller.schedules.go
│   │   ├── controller.todos.go
│   │   └── controller.users.go
│   ├── dbconfig.yml
│   ├── middleware
│   │   └── middleware.go
│   ├── migrations
│   │   └── development
│   ├── models
│   │   ├── model.common.go
│   │   ├── model.schedule.go
│   │   ├── model.user.go
│   │   └── model.user_test.go
│   └── routes.go
├── bin
│   ├── acmeprobe
│   ├── dlv
│   ├── godoc
│   ├── gore
│   └── sql-migrate
├── gin.log
├── go.mod
├── go.sum
├── main.go
├── note
│   ├── memo.md
│   ├── migration.md
│   ├── schedule_app_memo.md
│   └── test.md
└── pkg
    ├── mod
    │   ├── cache
    │   ├── github.com
    │   ├── go.starlark.net@v0.0.0-20200821142938-949cc6f4b097
    │   ├── golang.org
    │   ├── google.golang.org
    │   ├── gopkg.in
    │   └── gorm.io
    └── sumdb
        └── sum.golang.org

19 directories, 26 files
```
