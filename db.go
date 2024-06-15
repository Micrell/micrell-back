package main

import (
	"database/sql"
	"log"
)

type Storage interface {
	InsertProject(*Project) error
	DeleteProject(int) error
	UpdateProject(*Project) error
	GetProjects() ([]*Project, error)
	GetProjectByID(int) (*Project, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() *PostgresStore {
	connStr := "user=admin password=admin dbname=micrell-back sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresStore{db: db}
}

func (pg *PostgresStore) Init() error {
	err := pg.createProjectTable()
	return err
}

func (pg *PostgresStore) createProjectTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		url TEXT NOT NULL,
		last_mod_date DATE NOT NULL
	);`

	_, err := pg.db.Exec(query)
	return err
}

func (pg *PostgresStore) InsertProject(p *Project) error {
	query := `
	INSERT INTO projects 
	(title, url, last_mod_date) 
	VALUES ($1, $2, $3)`

	_, err := pg.db.Exec(query, p.Title, p.URL, p.LastModDate)
	return err
}

func DeleteProject(id int) error {
	return nil
}

func UpdateProject(p *Project) error {
	return nil
}

func (pg *PostgresStore) GetProjectByID(id int) (*Project, error) {
	query := `
	SELECT id, title, url, last_mod_date FROM projects WHERE id = $1
	`
	row := pg.db.QueryRow(query, id)
	project := new(Project)
	err := row.Scan(&project.ID, &project.Title, &project.URL, &project.LastModDate)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (pg *PostgresStore) GetProjects() ([]*Project, error) {
	query := `
	SELECT * FROM projects
	`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	projects := []*Project{}
	for rows.Next() {
		p := new(Project)
		err := rows.Scan(&p.ID, &p.Title, &p.URL, &p.LastModDate)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}
