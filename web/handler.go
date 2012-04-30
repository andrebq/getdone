package web

import (
	"fmt"
	"github.com/andrebq/getdone/entity"
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
		print("Fetching tasks by project Id: ", projId, "\n")
		lt.SelectProject(projId)
		tasks, err := lt.AllOpen()
		if err != nil {
			l.Error("Error while fetching task list: %v", err)
			http.Error(w, "Unable to fetch open tasks list", http.StatusInternalServerError)
		} else {
			_, err = WriteJson(w, tasksToJson(tasks), "", http.StatusOK)
			if err != nil {
				http.Error(w, "Unable to write the reponse", http.StatusInternalServerError)
			}
		}
	}
}

// Add a new task
func AddTask(w http.ResponseWriter, req *http.Request) {
	// think about this method later
	http.Error(w, "Not time to think about creating tasks", http.StatusInternalServerError)
}

func tasksToJson(t []*entity.Task) Json {
	print("Converting ", len(t), " tasks to json\n")
	j := make(Json)
	for _, task := range t {
		j.Push("tasks", jsonfyTask(task))
	}
	return j
}

func jsonfyTask(task *entity.Task) Json {
	return Json{"title": task.Title,
		"description": task.Description,
		"done":        task.Done,
		"id":          task.Id,
		"project": Json{
			"id":   task.Project.Id,
			"name": task.Project.Name,
		},
	}
}
