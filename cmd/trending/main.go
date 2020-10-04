package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ohatakky/ohatakkyp/functions/trending"
)

func init() {
	godotenv.Load()
}

func main() {
	err := trending.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
