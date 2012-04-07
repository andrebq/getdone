package uc

import (
	"github.com/andrebq/getdone/entity"
	"testing"
)

func TestCreateNewProject(t *testing.T) {
	cp := NewCreateProject()
	cp.ProjectRepo = new(MockProjectRepo)

	exp := &entity.Project{0, "test"}
	project, err := cp.Create(exp.Name)

	if err != nil {
		t.Fatalf("Unable to create project. Cause: %v", err)
	}

	if project.Id == 0 {
		t.Errorf("Project.Id shouldn't be 0")
	}

	if project.Name != exp.Name {
		t.Errorf("Expecting %v got %v", exp.Name, project.Name)
	}
}
