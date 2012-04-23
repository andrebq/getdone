package web

import (
	"net/http"
	"github.com/andrebq/getdone/uc"
	"github.com/andrebq/getdone/repo"
	"fmt"
	"io"
)

// API call for creating a new project
func CreateProject(w http.ResponseWriter, req *http.Request) {
	ctx := OpenCtx(req)
	session := Session(ctx)

	req.ParseForm()
	name := req.Form.Get("name")

	if name != "" {
		cp := uc.NewCreateProject()
		cp.ProjectRepo = repo.NewProject(session.DB("getdone"))
		p, err := cp.Create(name)
		if err != nil {
			http.Error(w, "Unable to create project repo", http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Location", fmt.Sprintf("%v", p.Id))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Required fields are empty (name)")
	}
}

// API call for listing all the pending tasks of a project
func ListTasks(w http.ResponseWriter, req *http.Request) {
	ctx := OpenCtx(req)
	session := Session(ctx)

	req.ParseForm()
	projName := req.URL.Query().Get("project")
	if projName == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Cannot list without project name")
	} else {
		lt := uc.NewListTasks()
		db := session.DB("getdone")
		lt.ProjectRepo = repo.NewProject(db)
		lt.TaskRepo = repo.NewTask(db, lt.ProjectRepo)
		open, err := lt.AllOpen()
		if err != nil {
			http.Error(w, "Unable to fetch open tasks list", http.StatusInternalServerError)
		} else {
			_, err = WriteJson(w, data, "", http.StatusOK)
			if err != nil {
				http.Error(w, "Unable to write the reponse", http.StatusInternalServerError)
			}
		}
	}
}
