package uc

import (
	"github.com/andrebq/getdone/entity"
	"reflect"
	"testing"
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

func TestCreateNewTask(t *testing.T) {
	ct := NewCreateTask()
	ct.ProjectRepo = new(MockProjectRepo)
	ct.TaskRepo = new(MockTaskRepo)
	project, err := ct.SelectProject("testproject")
	if err != nil {
		t.Fatalf("Unable to select project. Cause: %v", err)
	}

	et := &entity.Task{1, "This is the task title", "This is the task description", false, project}

	task, err := ct.CreateTask(et.Title, et.Description)
	if err != nil {
		t.Fatalf("Unable to create the task. Cause: %v", err)
	}

	if !reflect.DeepEqual(*task, *et) {
		t.Fatalf("Expecting %v got %v", et, task)
	}
}
