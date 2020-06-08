# 開発環境構築
```bash
$ docker-compose up -d
$ docker-compose exec server bash
$ cd client
$ yarn serve
```

## migration
```bash
$ go get -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate
$ migrate create -ext sql -dir migrations -seq create_users_table
$ migrate -source file://migrations/ -database 'mysql://root:root@tcp(mysql:3306)/accountant' up 1 # 指定
$ migrate -source file://migrations/ -database 'mysql://root:root@tcp(mysql:3306)/accountant' down # 全て
```