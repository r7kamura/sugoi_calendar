package main

import (
	"encoding/json"
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
	titles := make([]Title, len(records))
	for i, record := range records {
		titles[i] = *record.(*Title)
	}
	writeJsonResponse(writer, 200, titles)
}

func TitleCreateHandler(writer http.ResponseWriter, request *http.Request) {
	titleParam := request.FormValue("title")
	if titleParam == "" {
		writeJsonErrorResponse(writer, 400, "title parameter is required")
		return
	}
	title := &Title{Title: titleParam}
	err := dbMap.Insert(title)
	if err != nil {
		writeJsonErrorResponse(writer, 500, "Failed to insert a new title")
		return
	}
	writeJsonResponse(writer, 201, title)
}

func writeJsonResponse(writer http.ResponseWriter, statusCode int, data interface{}) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(json)
	return nil
}

func writeJsonErrorResponse(writer http.ResponseWriter, statusCode int, message string) error {
	return writeJsonResponse(writer, statusCode, NewErrorResponseBody(message))
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}

func NewErrorResponseBody(message string) ErrorResponseBody {
	return ErrorResponseBody{Message: message}
}
