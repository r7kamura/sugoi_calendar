package main

import (
	"encoding/json"
	"fmt"
	"github.com/r7kamura/router"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := router.NewRouter()
	router.Get("/titles", TitleIndexHandler)
	router.Post("/titles", TitleCreateHandler)
	http.ListenAndServe(":3000", router)
}

func TitleIndexHandler(writer http.ResponseWriter, request *http.Request) {
	records, err := dbMap.Select(Title{}, `SELECT * FROM titles ORDER BY id DESC`)
	if err != nil {
		writeJsonErrorResponse(writer, 500, "Failed to load titles from database")
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
		writeJsonErrorResponse(writer, 400, "title parameter is required")
		return
	}
	titleRecord := &Title{Title: titleParam}
	err := dbMap.Insert(titleRecord)
	if err != nil {
		writeJsonErrorResponse(writer, 500, "Failed to insert a new title")
		return
	}
	fmt.Fprintf(writer, "%t\n", titleRecord)
}

func writeJsonErrorResponse(writer http.ResponseWriter, statusCode int, message string) {
	writer.WriteHeader(statusCode)
	json, _ := json.Marshal(NewErrorResponseBody(message))
	writer.Write(json)
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}

func NewErrorResponseBody(message string) ErrorResponseBody {
	return ErrorResponseBody{Message: message}
}
