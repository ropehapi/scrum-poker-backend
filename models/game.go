package models

type Game struct {
	Id int `json:"id"`
	OwnerId int `json:"owner_id"`
	Name string `json:"name"`
}

var Games []Game