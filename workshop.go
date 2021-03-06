package main

import (
	"encoding/json"
	"fmt"
)

// Workshop is a type.
type Workshop struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"preview_url"`
	Level       int    `json:"level"`
}

// GetWorkshopInfo is a function that can get the infomation of a workshop.
func GetWorkshopInfo(id string) Workshop {
	pageBytes := Scrap("https://api.codemao.cn/web/shops/" + string(id))
	var workshop Workshop
	json.Unmarshal(pageBytes, &workshop)
	return workshop
}

// NotifyWorkshopInfo can print the info.
func NotifyWorkshopInfo(w Workshop) {
	fmt.Printf("Name:\t%s\n", w.Name)
	fmt.Println("Description: --------------------------------------")
	fmt.Printf("%s\n", w.Description)
	fmt.Println("---------------------------------------------------")
	fmt.Printf("Logo URL:\t%s\n", w.Logo)
	fmt.Printf("Level:\t%d\n", w.Level)
}
