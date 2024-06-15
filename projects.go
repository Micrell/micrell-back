package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v62/github"
	"github.com/joho/godotenv"
)

type Project struct {
	ID          int
	Title       string
	URL         string
	LastModDate time.Time
}

func NewProject(repo *github.Repository) *Project {
	return &Project{
		ID:          0,
		Title:       *repo.Name,
		URL:         *repo.HTMLURL,
		LastModDate: repo.UpdatedAt.Time,
	}
}

func (p Project) String() string {
	return fmt.Sprintf("{ID: %d, Title: %s, URL: %s, LasModDate: %s}",
		p.ID,
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
		projects = append(projects, *NewProject(repo))
	}
	return projects
}
