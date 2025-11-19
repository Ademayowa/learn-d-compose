package models

import (
	db "github.com/Ademayowa/learn-d-compose/internal/database"
	"github.com/google/uuid"
)

type Job struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// Save Jobs into the database
func (job *Job) Save() error {
	job.ID = uuid.New().String()

	query := `
		INSERT INTO jobs (id, title, description)
		VALUES ($1, $2, $3)
	`
	_, err := db.DB.Exec(query, job.ID, job.Title, job.Description)
	if err != nil {
		return err
	}

	return nil
}

// Retrieves all jobs from the database
func GetAll() ([]Job, error) {
	query := `
		SELECT * FROM jobs
		ORDER BY title ASC
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job

	for rows.Next() {
		var job Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, rows.Err()
}
