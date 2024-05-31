package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v62/github"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Project struct {
	ID          uuid.UUID
	Title       string
	URL         string
	LastModDate time.Time
}

func (p Project) String() string {
	return fmt.Sprintf("{ID: %s, Title: %s, URL: %s, LasModDate: %s}",
		p.ID.String(),
		p.Title,
		p.URL,
		p.LastModDate.Format("2006-01-02 15:04:05"))
}

func loadToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("GITHUB_TOKEN")
}

func downloadGithub() []Project {

	token := loadToken()
	client := github.NewClient(nil).WithAuthToken(token)
	opts := &github.RepositoryListByAuthenticatedUserOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	repos, _, err := client.Repositories.ListByAuthenticatedUser(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	var projects []Project
	for _, repo := range repos {
		projects = append(projects, Project{
			ID:          uuid.New(),
			Title:       *repo.Name,
			URL:         *repo.HTMLURL,
			LastModDate: repo.UpdatedAt.Time,
		})
	}
	return projects
}