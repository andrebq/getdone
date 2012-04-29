package web

import (
	"github.com/andrebq/getdone/repo"
	"github.com/andrebq/getdone/uc"
	"io"
	"net/http"
)

// API call for creating a new project
func CreateProject(w http.ResponseWriter, req *http.Request) {
	session := Session(req)

	req.ParseForm()
	name := req.Form.Get("name")

	if name != "" {
		cp := uc.NewCreateProject()
		cp.ProjectRepo = repo.NewProject(session.DB("getdone"))
		p, err := cp.Create(name)
		if err != nil {
			http.Error(w, "Unable to create project repo", http.StatusInternalServerError)
		} else {
			w.Header().Set("Location", ResolveRef(req, "../tasks", "projectid", int64(p.Id)).String())
			w.WriteHeader(http.StatusCreated)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Required fields are empty (name)")
	}
}

// API call for listing all the pending tasks of a project
func ListTasks(w http.ResponseWriter, req *http.Request) {
	session := Session(req)

	req.ParseForm()
	projName := req.URL.Query().Get("projectid")
	if projName == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Cannot list without project Id")
	} else {
		lt := uc.NewListTasks()
		db := session.DB("getdone")
		prepo := repo.NewProject(db)
		lt.ProjectRepo = prepo
		lt.TaskRepo = repo.NewTask(db, prepo)
		_, err := lt.AllOpen()
		if err != nil {
			http.Error(w, "Unable to fetch open tasks list", http.StatusInternalServerError)
		} else {
			_, err = WriteJson(w, nil, "", http.StatusOK)
			if err != nil {
				http.Error(w, "Unable to write the reponse", http.StatusInternalServerError)
			}
		}
	}
}
