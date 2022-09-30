package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	driverName = "mysql"
	dsn        = "migration-user:local-dev@tcp(localhost:3306)/local?multiStatements=true"
)

func main() {

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Print(err)
		return
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Print(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Print(err)
		return
	}

	err = m.Steps(-1)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("complete")

}
