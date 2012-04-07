package uc

import (
	"errors"
	"fmt"
	"github.com/andrebq/getdone/entity"
)

type MockProjectRepo struct {
	data map[int64]*entity.Project
}

func (m *MockProjectRepo) ensureData() {
	if m.data == nil {
		m.data = make(map[int64]*entity.Project)
	}
}

func (m *MockProjectRepo) ByName(name string) (*entity.Project, error) {
	m.ensureData()
	project := &entity.Project{0, name}
	m.Save(project)
	return project, nil
}

func (m *MockProjectRepo) Save(p *entity.Project) error {
	m.ensureData()
	if p.Id == 0 {
		p.Id = int64(len(m.data) + 1)
		m.data[p.Id] = p
	}
	m.data[p.Id] = p
	return nil
}

type MockTaskRepo struct {
	data map[int64]*entity.Task
}

func (m *MockTaskRepo) ensureData() {
	if m.data == nil {
		m.data = make(map[int64]*entity.Task)
	}
}

func (m *MockTaskRepo) Save(t *entity.Task) error {
	m.ensureData()
	if t.Id == 0 {
		t.Id = int64(len(m.data) + 1)
		m.data[t.Id] = t
	}
	m.data[t.Id] = t
	return nil
}

func (m *MockTaskRepo) ById(id int64) (*entity.Task, error) {
	m.ensureData()
	if t, has := m.data[id]; has {
		return t, nil
	}
	return nil, errors.New(fmt.Sprintf("Unable to find task %v", id))
}
