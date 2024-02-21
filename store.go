package main

import "database/sql"

type Store interface {
	// User Services
	CreateUser() error

	// Task services
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// User Services methods
func (s *Storage) CreateUser() error {
	return nil
}

// Task Services methods
func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec(
		`INSERT INTO tasks (name, status, peojectId, assignedToID)
		VALUES (?, ?, ?, ?)
		`, t.Name, t.Status, t.ProjectID, t.AssignedToID)
	if err != nil {
		return nil, err
	}

	// TODO: We can make a specific struct for createTask to avoid deal with ID special is create it and increment by DB itself
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow(`SELECT * FROM tasks WHERE id = ?`, id).Scan(
		&t.ID,
		&t.Name,
		&t.Status,
		&t.ProjectID,
		&t.AssignedToID,
		&t.CreatedAt,
	)

	return &t, err
}
