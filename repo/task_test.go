package repo

import (
	"github.com/andrebq/getdone/entity"
	"reflect"
	"testing"
)

func TestTaskRepo(t *testing.T) {
	p := &entity.Project{0, "test"}
	task := &entity.Task{0, "Do something", "Do something really interesting", false, p}

	session := openMgoSession(t)
	defer session.Close()

	db := session.DB("testdb")
	defer db.DropDatabase()
	pRepo := &Project{db}
	tRepo := &Task{db, pRepo}

	err := pRepo.Save(p)
	if err != nil {
		t.Fatalf("Unable to save project. Cause: %v", err)
	}

	// p now has a valid id
	err = tRepo.Save(task)
	if err != nil {
		t.Fatalf("Unable to save task. Cause %v", err)
	}

	if task.Id == 0 {
		t.Fatalf("The Task.Id property should have been updated.")
	}

	t2, err := tRepo.ById(task.Id)
	if err != nil {
		t.Fatalf("Error while fetching task by id (%v): Cause: %v", task.Id, err)
	}

	if !reflect.DeepEqual(task, t2) {
		t.Fatalf("Expecting %v got %v", task, t2)
	}
}

func TestAllByProj(t *testing.T) {
	p := &entity.Project{0, "test"}
	task := &entity.Task{0, "Do something", "Do something really interesting", false, p}

	session := openMgoSession(t)
	defer session.Close()

	db := session.DB("testdb")
	defer db.DropDatabase()
	pRepo := &Project{db}
	tRepo := &Task{db, pRepo}

	err := pRepo.Save(p)
	if err != nil {
		t.Fatalf("Unable to save project. Cause: %v", err)
	}

	// p now has a valid id
	err = tRepo.Save(task)
	if err != nil {
		t.Fatalf("Unable to save task. Cause %v", err)
	}

	if task.Id == 0 {
		t.Fatalf("The Task.Id property should have been updated.")
	}

	tasks, err := tRepo.AllByProject(p.Id)

	if err != nil {
		t.Fatalf("Error while fetching tasks by project (%v): Cause: %v", p.Id, err)
	}

	if tasks == nil || len(tasks) == 0 {
		t.Fatalf("Error while fetching tasks. Array should have 1 record.")
	}

	for _, t2 := range tasks {
		if !reflect.DeepEqual(task, t2) {
			t.Fatalf("Expecting %v got %v", task, t2)
		}
	}
}

func TestAllByState(t *testing.T) {
	p := &entity.Project{0, "test"}
	task := &entity.Task{0, "Do something", "Do something really interesting", false, p}

	session := openMgoSession(t)
	defer session.Close()

	db := session.DB("testdb")
	defer db.DropDatabase()
	pRepo := &Project{db}
	tRepo := &Task{db, pRepo}

	err := pRepo.Save(p)
	if err != nil {
		t.Fatalf("Unable to save project. Cause: %v", err)
	}

	// p now has a valid id
	err = tRepo.Save(task)
	if err != nil {
		t.Fatalf("Unable to save task. Cause %v", err)
	}

	if task.Id == 0 {
		t.Fatalf("The Task.Id property should have been updated.")
	}

	tasks, err := tRepo.AllByState(p.Id, false)

	if err != nil {
		t.Fatalf("Error while fetching tasks by project (%v): Cause: %v", p.Id, err)
	}

	if tasks == nil || len(tasks) == 0 {
		t.Fatalf("Error while fetching tasks. Array should have 1 record.")
	}

	for _, t2 := range tasks {
		if !reflect.DeepEqual(task, t2) {
			t.Fatalf("Expecting %v got %v", task, t2)
		}
	}
}
