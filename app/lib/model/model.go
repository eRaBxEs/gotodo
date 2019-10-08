package model

import (
	"time"

	"github.com/go-pg/pg"
)

// Task ...
type Task struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

// GetAll ...
func (s *Task) GetAll(db *pg.DB) ([]Task, error) {

	tasks := []Task{}

	if err := db.Model(&tasks).Select(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

// Save ...
func (s *Task) Save(db *pg.DB) error {

	if err := db.Insert(s); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (s *Task) Delete(db *pg.DB) error {

	if err := db.Delete(s); err != nil {
		return err
	}

	return nil
}
