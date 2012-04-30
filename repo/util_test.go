package repo

import (
	"launchpad.net/mgo"
	"log"
	"os"
	"testing"
)

func openMgoSession(t *testing.T) *mgo.Session {
	mgo.SetLogger(log.New(os.Stderr, "MONGO ", log.LstdFlags))
	mgo.SetDebug(true)
	s, err := mgo.Dial("localhost")
	if err != nil {
		t.Fatalf("Error opening mongo session. %v", err)
	}
	return s
}
