package main

import (
	"fmt"
)

func main() {
	projects := downloadGithub()

	for _, p := range projects {
		fmt.Println(p.String())
	}
}
