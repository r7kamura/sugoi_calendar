package main

import (
	"github.com/coopernurse/gorp"
	"time"
)

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
	UpdatedInSyobocalAt time.Time `db:"updated_in_syobocal_at" json:"updated_in_syobocal_at"`
}

func (title *Title) IsValid() bool {
	return title.Name != ""
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
