package main

import (
	"sia/config"
	"sia/routes"
)

func main() {
	dbConfig := config.DBConfig{
		Host:     "156.67.222.224",
		Port:     3306,
		User:     "u188337861_api",
		Password: "Mamangsekayu97",
		DBName:   "u188337861_api",
	}
	config.Init(dbConfig)

	routes.Handler()
}
