package uc

import (
	"github.com/andrebq/getdone/entity"
)

type CompleteTask struct {
	TaskRepo TaskRepo
	t        *entity.Task
}

func (c *CompleteTask) SelectTask(id int64) (*entity.Task, error) {
	var err error
	c.t, err = c.TaskRepo.ById(id)
	return c.t, err
}

func (c *CompleteTask) Complete() (*entity.Task, error) {
	c.t.Done = true
	err := c.TaskRepo.Save(c.t)
	return c.t, err
}

func NewCompleteTask() *CompleteTask {
	return new(CompleteTask)
}
