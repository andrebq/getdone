package web

import (
	"code.google.com/p/gorilla/context"
	"launchpad.net/mgo"
	"log"
	"net/http"
	"os"
)

type mongoSessionId int

var (
	session *mgo.Session
	sid     = mongoSessionId(0)
)

// Must initialize the mongo session
func InitMongo(url string) (err error) {
	mongolog := log.New(os.Stderr, "MONGO ", log.LstdFlags)
	mgo.SetLogger(mongolog)
	mgo.SetDebug(true)
	session, err = mgo.Dial(url)
	return err
}

// Clone a session for the given context
func Session(req *http.Request) *mgo.Session {
	if val := context.DefaultContext.Get(req, sid); val != nil {
		return val.(*mgo.Session)
	}
	return nil
}

// Ensure that a valid session is already present in the Context when hndl.ServeHTTP is called.
//
// The session is closed after the handler end's it's execution
func EnsureSession(hndl http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		s := Session(req)
		if s == nil {
			s = session.Copy()
			context.DefaultContext.Set(req, sid, s)
		}
		defer s.Close()
		defer context.DefaultContext.Delete(req, sid)
		hndl.ServeHTTP(w, req)
	})
}
