package uc

import (
	"github.com/andrebq/getdone/entity"
	"reflect"
	"testing"
)

func TestCreateNewTask(t *testing.T) {
	ct := NewCreateTask()
	ct.ProjectRepo = new(MockProjectRepo)
	ct.TaskRepo = new(MockTaskRepo)

	createPrj := NewCreateProject()
	createPrj.ProjectRepo = ct.ProjectRepo
	createPrj.Create("testproject")

	project, err := ct.SelectProject("testproject")
	if err != nil {
		t.Fatalf("Unable to select project. Cause: %v", err)
	}

	et := &entity.Task{1, "This is the task title", "This is the task description", false, project}

	task, err := ct.Create(et.Title, et.Description)
	if err != nil {
		t.Fatalf("Unable to create the task. Cause: %v", err)
	}

	if !reflect.DeepEqual(*task, *et) {
		t.Fatalf("Expecting %v got %v", et, task)
	}
}
