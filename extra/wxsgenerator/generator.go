package main

import (
	"log"
	"os"
	"text/template"

	"github.com/google/uuid"
)

func main() {
	templ, err := template.ParseFiles(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	templ.Execute(os.Stdout, uuid.New())
}
