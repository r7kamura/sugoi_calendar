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
	Abbreviation string `db:"abbreviation" json:"abbreviation"`
	CategoryID   int    `db:"category_id"  json:"category_id"`
	Comment      string `db:"comment"      json:"comment"`
	English      string `db:"english"      json:"english"`
	Hiragana     string `db:"hiragana"     json:"hiragana"`
	ID           int    `db:"id"           json:"id"`
	Name         string `db:"name"         json:"name"`
	UpdatedAt    string `db:"updated_at"   json:"updated_at"`
}
