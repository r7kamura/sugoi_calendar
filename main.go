package main

import (
	"github.com/r7kamura/router"
	"net/http"
)

var sugoiCalendarHandler *router.Router

func init() {
	sugoiCalendarHandler = router.NewRouter()
	sugoiCalendarHandler.Get("/titles", TitleIndexHandler)
	sugoiCalendarHandler.Get("/titles/:id", TitleShowHandler)
	sugoiCalendarHandler.Post("/titles", TitleCreateHandler)
}

func main() {
	http.ListenAndServe(":3000", sugoiCalendarHandler)
}

func TitleIndexHandler(writer http.ResponseWriter, request *http.Request) {
	TitleController{
		Controller{
			ResponseWriter: writer,
			Request: request,
		},
	}.Index()
}

func TitleShowHandler(writer http.ResponseWriter, request *http.Request) {
	TitleController{
		Controller{
			ResponseWriter: writer,
			Request: request,
		},
	}.Show()
}

func TitleCreateHandler(writer http.ResponseWriter, request *http.Request) {
	TitleController{
		Controller{
			ResponseWriter: writer,
			Request: request,
		},
	}.Create()
}
