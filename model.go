package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var dbMap *gorp.DbMap

const (
	dataSourceName    string = "root:@unix(/tmp/mysql.sock)/sugoi_calendar_development?parseTime=true&loc=Asia/Tokyo"
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
	Abbreviation        string    `db:"abbreviation"           json:"abbreviation"`
	CategoryID          int       `db:"category_id"            json:"category_id"`
	Comment             string    `db:"comment"                json:"comment"`
	CreatedAt           time.Time `db:"created_at"             json:"created_at"`
	English             string    `db:"english"                json:"english"`
	Hiragana            string    `db:"hiragana"               json:"hiragana"`
	ID                  int       `db:"id"                     json:"id"`
	Name                string    `db:"name"                   json:"name"`
	UpdatedAt           time.Time `db:"updated_at"             json:"updated_at"`
	UpdatedInSyobocalAt string    `db:"updated_in_syobocal_at" json:"updated_in_syobocal_at"`
}

func (title *Title) PreInsert(executor gorp.SqlExecutor) error {
	title.CreatedAt = clock.Now()
	title.UpdatedAt = title.CreatedAt
	return nil
}

func (title *Title) PreUpdated(executor gorp.SqlExecutor) error {
	title.UpdatedAt = clock.Now()
	return nil
}
