package web

import (
	"fmt"
	"github.com/andrebq/getdone/repo"
	"github.com/andrebq/getdone/uc"
	"io"
	"net/http"
	"strconv"
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
	projId, err := strconv.ParseInt(req.URL.Query().Get("projectid"), 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot list without a valid project id. Error: %v", err), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		lt := uc.NewListTasks()
		db := session.DB("getdone")
		prepo := repo.NewProject(db)
		lt.ProjectRepo = prepo
		lt.TaskRepo = repo.NewTask(db, prepo)
		lt.SelectProject(projId)
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

// Add a new task
func AddTask(w http.ResponseWriter, req *http.Request) {
	session := Session(req)

	req.ParseForm()
	projName := req.Form.Get("projectid")
	taskTitle := req.Form.Get("title")
	if projName == "" || taskTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Cannot insert a task without an projectid or a title")
	} else {
		ct := uc.NewCreateTask()
		db := session.DB("getdone")
		prepo := repo.NewProject(db)
		ct.ProjectRepo = prepo
		ct.TaskRepo = repo.NewTask(db, prepo)
		_, err := ct.Create(taskTitle, "")
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
