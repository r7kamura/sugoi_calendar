package main

import (
	"fmt"
	"github.com/r7kamura/router"
	"net/http"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := router.NewRouter()
	router.Get("/titles", http.HandlerFunc(TitleIndexHandler))
	router.Post("/titles", http.HandlerFunc(TitleCreateHandler))
	http.ListenAndServe(":3000", router)
}

func TitleIndexHandler(writer http.ResponseWriter, request *http.Request) {
	records, err := dbMap.Select(Title{}, `SELECT * FROM titles ORDER BY id DESC`)
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

func TitleCreateHandler(writer http.ResponseWriter, request *http.Request) {
	titleParam := request.FormValue("title")
	if titleParam == "" {
		writer.WriteHeader(400)
		return
	}
	idParam := request.FormValue("id")
	if idParam == "" {
		writer.WriteHeader(400)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writer.WriteHeader(400)
		return
	}
	titleRecord := &Title{
		ID: id,
		Title: titleParam,
	}
	err = dbMap.Insert(titleRecord)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	fmt.Fprintf(writer, "%t\n", titleRecord)
}
