package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"testing"
	"time"
)

var (
	mgoStarted bool = false
)

func initMgo(t *testing.T) {
	if mgoStarted {
		return
	}

	err := InitMongo("localhost")
	if err != nil {
		t.Fatalf("Unable to start mgo session. %v", err)
	}
	mgoStarted = true
}

func insertInMgo(t *testing.T, colName string, doc Json) {
	err := session.DB("getdone").C(colName).Insert(doc)
	if err != nil {
		t.Fatalf("Error saving %v to mgo. Cause: %v", doc, err)
	}
}

func readResponseBody(t *testing.T, resp *http.Response) string {
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unable to read the response body. Cause: %v", err)
	}
	return string(contents)
}

func TestCreateProject(t *testing.T) {
	initMgo(t)
	server := httptest.NewServer(EnsureSession(http.HandlerFunc(CreateProject)))
	defer server.Close()

	resp, err := http.PostForm(server.URL, url.Values{"name": {"testproject"}})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	readResponseBody(t, resp)
	loc := resp.Header.Get("Location")
	locre := regexp.MustCompile("/tasks\\?projectid=\\d+")
	if !locre.MatchString(loc) {
		t.Errorf("Location %v should match against %v", loc, locre)
	}
}

func TestListTasks(t *testing.T) {
	initMgo(t)
	server := httptest.NewServer(EnsureSession(http.HandlerFunc(ListTasks)))

	s := session.Copy()
	defer s.Close()

	proj := make(Json).Put("name", "prjtest").Put("id", time.Now().UnixNano())
	insertInMgo(t, "projects", proj)
	task := make(Json).Put("title", "task1").Put("description", "task1").Put("done", false).Put("projectid", proj.Get("id").(int64)).Put("id", time.Now().UnixNano())
	insertInMgo(t, "tasks", task)

	u, _ := url.Parse(server.URL)
	q := u.Query()
	q.Set("projectid", strconv.FormatInt(proj.Get("id").(int64), 10))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	body := readResponseBody(t, resp)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid status code. Expected %v got %v (%v)", http.StatusOK, resp.StatusCode, resp.Status)
	}

	respJson := make(Json)
	err = json.Unmarshal([]byte(body), &respJson)
	if err != nil {
		t.Fatalf("Error while reading the json object from response body. %v", err)
	}

	// remove the projectid property and put the project object
	// since the api returns the project inside the task object
	delete(task, "projectid")
	task.Put("project", proj)
	respJson = Json(respJson.Array("tasks")[0].(map[string]interface{}))
	// problem with longs
	// the encoding/json package converts numbers to float64 when using map[string]interface{}
	// so, this code unmarshal a Json into itself to force the float64 conversion.
	// no error should happen since the Json object is valid
	json.Unmarshal([]byte(task.String()), &task)
	if !reflect.DeepEqual(task, respJson) {
		t.Errorf("Expecting %v got %v", task, respJson)
	}
}
