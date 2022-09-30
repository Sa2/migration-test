# setup

```
$ go get github.com/golang-migrate/migrate/v4
$ go mod download github.com/go-sql-driver/mysql

$ brew install golang-migrate
```



# create migrate files

```bash
$ migrate create -ext sql -dir migrations -seq create_table1
```