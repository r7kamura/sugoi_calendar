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

func BeReadableLike(values ...interface{}) string {
	withoutLineBreaks := strings.Replace(values[1].(string), "\n", "", -1)
	withoutTabs := strings.Replace(withoutLineBreaks, "\t", "", -1)
	return BeReadableAs(values[0], withoutTabs)
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
			dbMap.Insert(&Title{Name: "test"})
		})

		It("returns titles as JSON", func() {
			response, _ := http.Get(server.URL + "/titles")
			Expect(response.StatusCode).To(Equal, 200)
			Expect(response.Header).To(HaveJSONContentType)
			Expect(response.Body).To(BeReadableLike, `
				[
					{
						"abbreviation":"",
						"category_id":0,
						"comment":"",
						"english":"",
						"hiragana":"",
						"id":1,
						"name":"test",
						"updated_at":""
					}
				]`,
			)
		})
	})

	Describe(t, "GET /titles/:id", func() {
		Before(func() {
			dbMap.DropTablesIfExists()
			dbMap.CreateTables()
			dbMap.Insert(&Title{Name: "test"})
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
				Expect(response.Body).To(BeReadableLike, `
					{
						"abbreviation":"",
						"category_id":0,
						"comment":"",
						"english":"",
						"hiragana":"",
						"id":1,
						"name":"test",
						"updated_at":""
					}`,
				)
			})
		})
	})

	Describe(t, "POST /titles", func() {
		Before(func() {
			dbMap.DropTablesIfExists()
			dbMap.CreateTables()
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
				response, _ := http.PostForm(server.URL + "/titles", url.Values{"name": []string{"test"}})
				Expect(response.StatusCode).To(Equal, 406)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableAs, `{"message":"Request body must be a JSON encoded value"}`)
			})
		})

		Context("with JSON encoded request body", func() {
			It("creates a new title record & returns it", func() {
				response, _ := http.Post(server.URL + "/titles", "application/json", strings.NewReader(`{"name":"test"}`))
				Expect(response.StatusCode).To(Equal, 201)
				Expect(response.Header).To(HaveJSONContentType)
				Expect(response.Body).To(BeReadableLike, `
					{
						"abbreviation":"",
						"category_id":0,
						"comment":"",
						"english":"",
						"hiragana":"",
						"id":1,
						"name":"test",
						"updated_at":""
					}`,
				)
			})
		})
	})
}
