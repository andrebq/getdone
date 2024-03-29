package repo

import (
	"github.com/andrebq/getdone/entity"
	"reflect"
	"testing"
)

func TestProject(t *testing.T) {
	p := &entity.Project{0, "test"}

	session := openMgoSession(t)
	defer session.Close()

	repo := &Project{session.DB("testdb")}
	defer repo.db.DropDatabase()

	err := repo.Save(p)
	if err != nil {
		t.Fatalf("Error while saving project. %v", err)
	}

	p2, err := repo.ById(p.Id)
	if err != nil {
		t.Fatalf("Error while fetching project. %v", err)
	}

	if !reflect.DeepEqual(p2, p) {
		t.Errorf("Expecting %v got %v", p2, p)
	}
}
