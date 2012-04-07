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

func (l *ListTasks) SelectProject(name string) (*entity.Project, error) {
	var err error
	l.p, err = l.ProjectRepo.ByName(name)
	return l.p, err
}

// Return all the tasks that are currently open
func (l *ListTasks) AllOpen() ([]*entity.Task, error) {
	return l.TaskRepo.AllByState(l.p.Id, false)
}

func NewListTasks() *ListTasks {
	return &ListTasks{}
}
