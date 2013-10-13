package main

import (
	"fmt"
	"github.com/r7kamura/router"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := router.NewRouter()
	router.Get("/titles", http.HandlerFunc(TitleIndexHandler))
	http.ListenAndServe(":3000", router)
}

func TitleIndexHandler(writer http.ResponseWriter, request *http.Request) {
	records, err := dbMap.Select(Title{}, `SELECT * FROM titles ORDER BY id DESC LIMIT ?`, 1)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	titles := []Title{}
	for _, record := range records {
		titles = append(titles, *record.(*Title))
	}
	fmt.Fprintln(
		writer,
		titles,
	)
}
