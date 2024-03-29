package main

import (
	"database/sql"
)

type Store interface {
	// Project services
	CreateProject(p *Project) (*Project, error)
	GetProject(id string) (*Project, error)
	DeleteProject(id string) error

	// User Services
	CreateUser(u *User) (*User, error)
	LoginUser(email string) (*UserLoginRequest, error)
	GetUserByID(id string) (*User, error)

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

func (s *Storage) CreateProject(p *Project) (*Project, error) {
	rows, err := s.db.Exec(
		`INSERT INTO projects (name)
		VALUES (?)
		`, p.Name)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	p.ID = id
	return p, nil
}

func (s *Storage) GetProject(id string) (*Project, error) {
	var p Project
	err := s.db.QueryRow(`SELECT * FROM projects WHERE id = ?`, id).Scan(
		&p.ID,
		&p.Name,
		&p.CreatedAt,
	)

	return &p, err
}

func (s *Storage) DeleteProject(id string) error {
	_, err := s.db.Exec(`DELETE FROM projects WHERE id = ?`, id)
	return err
}

// User Services methods
func (s *Storage) CreateUser(u *User) (*User, error) {
	rows, err := s.db.Exec(
		`INSERT INTO users (email, firstName, lastName, password)
		VALUES (?, ?, ?, ?)
		`, u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (s *Storage) LoginUser(email string) (*UserLoginRequest, error) {
	var u UserLoginRequest
	err := s.db.QueryRow(`SELECT id, email, password FROM users WHERE email = ?`, email).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
	)

	return &u, err
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

func (s *Storage) GetUserByID(id string) (*User, error) {
	var u User
	err := s.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id).Scan(
		&u.ID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.Password,
		&u.CreatedAt,
	)

	return &u, err
}
