package uc

import (
	"testing"
)

func TestListOpenTasks(t *testing.T) {
	tRepo := new(MockTaskRepo)
	pRepo := new(MockProjectRepo)

	createPrj := NewCreateProject()
	createPrj.ProjectRepo = pRepo
	createPrj.Create("test")

	list := NewListTasks()
	list.TaskRepo = tRepo
	list.ProjectRepo = pRepo
	list.SelectProject("test")

	complete := NewCompleteTask()
	complete.TaskRepo = list.TaskRepo

	ct := NewCreateTask()
	ct.TaskRepo = list.TaskRepo
	ct.ProjectRepo = list.ProjectRepo
	ct.SelectProject("test")

	pending, _ := ct.Create("task 1", "Execute task 1")

	done, err := ct.Create("task 2", "Execute task 2")
	complete.SelectTask(done.Id)
	complete.Complete()

	tasks, _ := list.AllOpen()

	if err != nil {
		t.Fatalf("Unable to fetch the task list. Cause: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expecting %v tasks, got %v", 1, len(tasks))
	}

	if tasks[0].Id != pending.Id {
		t.Errorf("Pending task expected %v but got %v", pending.Id, tasks[0].Id)
	}
}

func TestListAllTasks(t *testing.T) {
	tRepo := new(MockTaskRepo)
	pRepo := new(MockProjectRepo)

	createPrj := NewCreateProject()
	createPrj.ProjectRepo = pRepo
	createPrj.Create("test")

	list := NewListTasks()
	list.TaskRepo = tRepo
	list.ProjectRepo = pRepo
	list.SelectProject("test")

	complete := NewCompleteTask()
	complete.TaskRepo = list.TaskRepo

	ct := NewCreateTask()
	ct.TaskRepo = list.TaskRepo
	ct.ProjectRepo = list.ProjectRepo
	ct.SelectProject("test")

	ct.Create("task 1", "Execute task 1")

	done, err := ct.Create("task 2", "Execute task 2")
	complete.SelectTask(done.Id)
	complete.Complete()

	// this should return all tasks
	tasks, _ := list.All()

	if err != nil {
		t.Fatalf("Unable to fetch the task list. Cause: %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf("Expecting %v tasks, got %v", 2, len(tasks))
	}
}
