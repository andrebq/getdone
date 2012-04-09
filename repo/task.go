package repo

import (
	"launchpad.net/mgo"
	"github.com/andrebq/getdone/entity"
	"github.com/andrebq/getdone/data"
	"time"
	"launchpad.net/mgo/bson"
)

type Task struct {
	db *mgo.Database
}

func (t *Task) Save(task *entity.Task) error {
	if task.Id == 0 {
		task.Id = time.Now().UnixNano()
	}
	dt := &data.Task{task.Id, task.Title, task.Description, task.Done, task.Project.Id}
	
	_, err := t.db.C("tasks").Upsert(bson.M{"id": task.Id}, dt)
	return err
}

func (t *Task) ById(id int64, pRepo *Project) (*entity.Task, error) {
	dt := &data.Task{}
	
	err := t.db.C("tasks").Find(bson.M{"id": id},).One(&dt)
	
	if err != nil {
		return nil, err
	}
	
	ret := &entity.Task{dt.Id, dt.Title, dt.Description, dt.Done, nil}
	
	p, err := pRepo.ById(dt.ProjectId)
	if err != nil {
		return nil, err
	}
	ret.Project = p
	
	return ret, nil
}
