package repo

import (
	"github.com/andrebq/getdone/data"
	"github.com/andrebq/getdone/entity"
	"launchpad.net/mgo"
	"launchpad.net/mgo/bson"
	"time"
)

// Responsible for storing project objects
type Project struct {
	db *mgo.Database
}

// Save the project in the database
func (p *Project) Save(proj *entity.Project) error {

	if proj.Id == 0 {
		proj.Id = time.Now().UnixNano()
	}

	dt := &data.Project{proj.Id, proj.Name}
	_, err := p.db.C("projects").Upsert(bson.M{"Id": proj.Id}, dt)
	return err
}

// Fetch the project by name
func (p *Project) ByName(name string) (*entity.Project, error) {
	dt := &data.Project{}
	err := p.db.C("projects").Find(bson.M{"name": name}).One(&dt)
	return &entity.Project{dt.Id, dt.Name}, err
}
