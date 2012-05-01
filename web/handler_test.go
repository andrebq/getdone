package web

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"io/ioutil"
	"regexp"
)

var (
	mgoStarted bool = false

)

func initMgo(t *testing.T) {
	if mgoStarted { return }

	err := InitMongo("localhost")
	if err != nil {
		t.Fatalf("Unable to start mgo session. %v", err)
	}
	mgoStarted = true
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

	resp, err := http.PostForm(server.URL, url.Values{"name":{"testproject"}})
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
