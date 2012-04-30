package uc

import (
	"github.com/andrebq/getdone/entity"
)

// Repository to work with projects
type ProjectRepo interface {
	ById(id int64) (*entity.Project, error)
	Save(p *entity.Project) error
}

// Repository to work with tasks
type TaskRepo interface {
	Save(t *entity.Task) error
	ById(id int64) (*entity.Task, error)
	AllByState(projId int64, done bool) ([]*entity.Task, error)
	AllByProject(projId int64) ([]*entity.Task, error)
}
