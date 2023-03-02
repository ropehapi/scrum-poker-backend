package models

const HOST = "HOST"
const GUEST = "GUEST"

type Player struct {
	Id       int `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Players []Player
