package main

type MockStore struct{}

// Project services
func (m *MockStore) CreateProject(p *Project) (*Project, error) {
	return &Project{}, nil
}

func (m *MockStore) GetProject(id string) (*Project, error) {
	return &Project{}, nil
}

func (m *MockStore) DeleteProject(id string) error {
	return nil
}

func (m *MockStore) CreateUser() error {
	return nil
}

func (m *MockStore) GetUserByID(id string) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}

func (m *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}
