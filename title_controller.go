package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type TitleController struct {
	Controller
}

func (controller TitleController) Index() {
	records, err := dbMap.Select(Title{}, `SELECT * FROM titles ORDER BY id DESC`)
	if err != nil {
		fmt.Println(err.Error())
		controller.RenderErrorJson(500, "Failed to load titles from database")
		return
	}
	titles := make([]Title, len(records))
	for i, record := range records {
		titles[i] = *record.(*Title)
	}
	controller.RenderJson(200, titles)
}

func (controller TitleController) Show() {
	id, err := strconv.Atoi(controller.Request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		controller.RenderErrorJson(400, "id parameter must be a positive integer")
		return
	}
	record, err := dbMap.Get(Title{}, id)
	if err != nil {
		fmt.Println(err.Error())
		controller.RenderErrorJson(500, "Failed to load record from database")
		return
	}
	if record == nil {
		controller.RenderErrorJson(404, "Not found")
		return
	}
	controller.RenderJson(200, record)
}

func (controller TitleController) Create() {
	var givenTitle GivenTitle
	decoder := json.NewDecoder(controller.Request.Body)
	err := decoder.Decode(&givenTitle)
	if err != nil {
		controller.RenderErrorJson(406, "Request body must be a JSON encoded value")
		return
	}
	if givenTitle.Name == "" {
		controller.RenderErrorJson(400, "title parameter is required")
		return
	}
	title := &Title{Name: givenTitle.Name}
	err = dbMap.Insert(title)
	if err != nil {
		controller.RenderErrorJson(500, "Failed to insert a new title")
		return
	}
	controller.RenderJson(201, title)
}

type GivenTitle struct {
	Abbreviation        string `json:"abbreviation"`
	CategoryID          int    `json:"category_id"`
	Comment             string `json:"comment"`
	English             string `json:"english"`
	Hiragana            string `json:"hiragana"`
	Name                string `json:"name"`
	UpdatedInSyobocalAt string `json:"updated_in_syobocal_at"`
}
