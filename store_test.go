package main

//Mocks

type MockStore struct {
}

func (m *MockStore) CreateProject(p *Project) error {
	return nil
}

func (m *MockStore) GetProject(id string) error {
	return nil
}

func (m *MockStore) DeleteProject(id string) error {
	return nil
}

func (m *MockStore) CreateUser() (*User, error) {
	return nil, nil
}

func (m *MockStore) GetUserByID(id string) (*User, error) {
	return nil, nil
}

func (m *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}

func (m *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}
