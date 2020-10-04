package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ohatakky/ohatakkyp/functions/blog"
)

func init() {
	godotenv.Load()
}

func main() {
	err := blog.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
