package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	pg := NewPostgresStore()
	err := pg.Init()
	if err != nil {
		log.Fatal("Error while init the db")
	}

	projects := downloadGithub()

	for _, p := range projects {
		fmt.Println(p.String())
	}
}
