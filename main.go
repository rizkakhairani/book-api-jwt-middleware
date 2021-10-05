package main

import (
	"book-api-jwt/config"
	"book-api-jwt/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}