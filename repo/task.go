package repo

import (
	"github.com/andrebq/getdone/data"
	"github.com/andrebq/getdone/entity"
	"launchpad.net/mgo"
	"launchpad.net/mgo/bson"
	"time"
)

type Task struct {
	db      *mgo.Database
	Project *Project
}

func (t *Task) Save(task *entity.Task) error {
	if task.Id == 0 {
		task.Id = time.Now().UnixNano()
	}
	dt := &data.Task{task.Id, task.Title, task.Description, task.Done, task.Project.Id}

	_, err := t.db.C("tasks").Upsert(bson.M{"id": task.Id}, dt)
	return err
}

func (t *Task) ById(id int64) (*entity.Task, error) {
	dt := &data.Task{}

	err := t.db.C("tasks").Find(bson.M{"id": id}).One(&dt)

	if err != nil {
		return nil, err
	}

	ret := &entity.Task{dt.Id, dt.Title, dt.Description, dt.Done, nil}

	p, err := t.Project.ById(dt.ProjectId)
	if err != nil {
		return nil, err
	}
	ret.Project = p

	return ret, nil
}

func NewTask(db *mgo.Database, pRepo *Project) *Task {
	return &Task{db, pRepo}
}
