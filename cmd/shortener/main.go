package main

import (
	"log"

	"github.com/learies/go-shortener/internal/app"
)

func main() {
	application := app.New()

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
