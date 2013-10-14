package main

import (
	. "github.com/r7kamura/gospel"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func BeReadableAs(values ...interface{}) string {
	data, _ := ioutil.ReadAll(values[0].(io.Reader))
	return Equal(string(data), values[1])
}

func HaveJSONContentType(values ...interface{}) string {
	return Equal(values[0].(http.Header).Get("Content-Type"), "application/json; charset=utf-8")
}

func TestSugoiCalendarHandler(t *testing.T) {
	server := httptest.NewServer(sugoiCalendarHandler)
	defer server.Close()

	Describe(t, "GET /titles", func() {
		Before(func() {
			dbMap.DropTablesIfExists()
			dbMap.CreateTablesIfNotExists()
			dbMap.Insert(&Title{Title: "testTitle"})
		})


		It("returns titles as JSON", func() {
			response, _ := http.Get(server.URL + "/titles")
			Expect(response.StatusCode).To(Equal, 200)
			Expect(response.Header).To(HaveJSONContentType)
			Expect(response.Body).To(BeReadableAs, `[{"ID":1,"Title":"testTitle"}]`)
		})
	})

	Describe(t, "GET /titles/:id", func() {
		Before(func() {
			dbMap.DropTablesIfExists()
			dbMap.CreateTables()
			dbMap.Insert(&Title{Title: "testTitle"})
		})

		Context("with non-integer id", func() {
			It("returns 400 error", func() {
				response, _ := http.Get(server.URL + "/titles/a")
				Expect(response.StatusCode).To(Equal, 400)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"message":"id parameter must be a positive integer"}`)
			})
		})

		Context("with negative integer id", func() {
			It("returns 400 error", func() {
				response, _ := http.Get(server.URL + "/titles/-1")
				Expect(response.StatusCode).To(Equal, 400)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"message":"id parameter must be a positive integer"}`)
			})
		})

		Context("with non-existent record's id", func() {
			It("returns 404 error", func() {
				response, _ := http.Get(server.URL + "/titles/2")
				Expect(response.StatusCode).To(Equal, 404)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"message":"Not found"}`)
			})
		})

		Context("with existent record's id", func() {
			It("returns the record", func() {
				response, _ := http.Get(server.URL + "/titles/1")
				Expect(response.StatusCode).To(Equal, 200)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"ID":1,"Title":"testTitle"}`)
			})
		})
	})

	Describe(t, "POST /titles", func() {
		Before(func() {
			dbMap.DropTablesIfExists()
			dbMap.CreateTables()
		})

		Context("with JSON encoded request body", func() {
			It("creates a new title record & returns it", func() {
				response, _ := http.Post(server.URL + "/titles", "application/json", strings.NewReader(`{"title":"testTitle"}`))
				Expect(response.StatusCode).To(Equal, 201)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"ID":1,"Title":"testTitle"}`)
			})
		})

		Context("without title parameter", func() {
			It("returns 400 error", func() {
				response, _ := http.Post(server.URL + "/titles", "application/json", strings.NewReader(`{}`))
				Expect(response.StatusCode).To(Equal, 400)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"message":"title parameter is required"}`)
			})
		})

		Context("with URL encoded request body", func() {
			It("returns 406 error", func() {
				response, _ := http.PostForm(server.URL + "/titles", url.Values{"title": []string{"testTitle"}})
				Expect(response.StatusCode).To(Equal, 406)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"message":"Request body must be a JSON encoded value"}`)
			})
		})
	})
}
