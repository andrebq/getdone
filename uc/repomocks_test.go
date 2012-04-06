package uc

import (
	"github.com/andrebq/getdone/entity"
)

type MockProjectRepo struct{}

func (m MockProjectRepo) ByName(name string) (*entity.Project, error) {
	return &entity.Project{1, name}, nil
}

type MockTaskRepo struct{}

func (m MockTaskRepo) Save(t *entity.Task) error {
	t.Id = 1
	return nil
}
