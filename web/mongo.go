package web

import (
	"launchpad.net/mgo"
	"net/http"
)

type mongoSessionId int

var (
	session *mgo.Session
	sid = mongoSessionId(0)
)

// Must initialize the mongo session
func InitMongo(url string) ( err error ){
	session, err = mgo.Dial(url)
	return err
}

// Clone a session for the given context
func Session(ctx *Context) *mgo.Session {
	if s, has := ctx.Get(sid); has {
		return s.(*mgo.Session)
	}
	s := session.Clone()
	ctx.Set(sid, s)
	return s
}

// Ensure that a valid session is already present in the Context when hndl.ServeHTTP is called.
//
// The session is closed after the handler end's it's execution
func EnsureSession(hndl http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		ctx := OpenCtx(req)
		s := Session(ctx)
		defer s.Close()
		hndl.ServeHTTP(w, req)
	})
}
