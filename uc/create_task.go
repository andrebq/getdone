package uc

import (
	"github.com/andrebq/getdone/entity"
)

// Handle the creating of new tasks for a given project
type CreateTask struct {
	ProjectRepo ProjectRepo
	TaskRepo    TaskRepo
	project     *entity.Project
}

// Select the project to create the task
func (c *CreateTask) SelectProject(id int64) (*entity.Project, error) {
	var err error
	c.project, err = c.ProjectRepo.ById(id)

	return c.project, err
}

func (c *CreateTask) Create(name, description string) (*entity.Task, error) {
	if c.project == nil {
		return nil, newInvalidState("The project is required to create a new task. Use Select")
	}

	t := &entity.Task{0, name, description, false, c.project}
	err := c.TaskRepo.Save(t)
	return t, err
}

func NewCreateTask() *CreateTask {
	return new(CreateTask)
}
