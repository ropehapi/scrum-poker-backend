package main

import (
	"github.com/ropehapi/scrum-poker/database"
	"github.com/ropehapi/scrum-poker/routes"
)

func main() {
	database.GetConn()
	routes.HandleRequests()
}
