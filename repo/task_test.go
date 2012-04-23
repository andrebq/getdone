package repo

import (
	"github.com/andrebq/getdone/entity"
	"launchpad.net/mgo"
	"reflect"
	"testing"
)

func TestTaskRepo(t *testing.T) {
	p := &entity.Project{0, "test"}
	task := &entity.Task{0, "Do something", "Do something really interesting", false, p}

	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Fatalf("Error while connecting to mongo. %v", err)
	}
	defer session.Close()

	db := session.DB("testdb")
	defer db.DropDatabase()
	pRepo := &Project{db}
	tRepo := &Task{db, pRepo}

	err = pRepo.Save(p)
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