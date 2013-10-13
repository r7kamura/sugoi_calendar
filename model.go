package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var dbMap *gorp.DbMap

const (
	dataSourceName    string = "root:@unix(/tmp/mysql.sock)/sugoi_calendar_development"
	tableType         string = "InnoDB"
	tableCharacterSet string = "UTF8"
)

func init() {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{tableType, tableCharacterSet}}
	dbMap.AddTableWithName(Title{}, "titles").SetKeys(true, "id")
}

type Title struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
