package uc

import (
	"github.com/andrebq/getdone/entity"
)

// Used to create new projects
type CreateProject struct {
	ProjectRepo ProjectRepo
	p           *entity.Project
}

func NewCreateProject() *CreateProject {
	return &CreateProject{}
}

func (c *CreateProject) Create(name string) (*entity.Project, error) {
	c.p = &entity.Project{0, name}
	err := c.ProjectRepo.Save(c.p)
	return c.p, err
}
