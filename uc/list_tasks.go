package uc

import (
	"github.com/andrebq/getdone/entity"
)

// Access the lists of tasks for a project.
type ListTasks struct {
	ProjectRepo ProjectRepo
	TaskRepo    TaskRepo
	p           *entity.Project
}

func (l *ListTasks) SelectProject(id int64) (*entity.Project, error) {
	var err error
	l.p, err = l.ProjectRepo.ById(id)
	return l.p, err
}

// Return all the tasks that are currently open
func (l *ListTasks) AllOpen() ([]*entity.Task, error) {
	return l.TaskRepo.AllByState(l.p.Id, false)
}

// Return all the tasks both open and closed
func (l *ListTasks) All() ([]*entity.Task, error) {
	return l.TaskRepo.AllByProject(l.p.Id)
}

func NewListTasks() *ListTasks {
	return &ListTasks{}
}
