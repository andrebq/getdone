package uc

import (
	"testing"
)

func TestCompleteTask(t *testing.T) {

	create := NewCreateTask()
	create.ProjectRepo = new(MockProjectRepo)
	create.TaskRepo = new(MockTaskRepo)

	create.SelectProject("test")
	create.CreateTask("title", "desc")

	ct := NewCompleteTask()
	ct.TaskRepo = create.TaskRepo

	task, err := ct.SelectTask(1)

	if err != nil {
		t.Fatalf("Unable to find a task. Cause: %v", err)
	}

	if task == nil {
		t.Fatalf("Task shouldn't be nil")
	}

	task, err = ct.Complete()
	if err != nil {
		t.Fatalf("Unable to complete the task. Cause: %v", err)
	}

	if !task.Done {
		t.Fatalf("Task should have been marked as done.")
	}
}
