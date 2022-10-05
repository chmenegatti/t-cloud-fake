package main

import (
	"github.com/chmenegatti/t-cloud-fake/database"
	"github.com/chmenegatti/t-cloud-fake/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
