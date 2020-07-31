// cmd/api/main.go

package main

import (
	"log"

	"github.com/go.mux/pkg/application"
	"github.com/go.mux/pkg/exithandler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}