package web

import (
	"sync"
	"net/http"
)

// Store random data.
//
// Items are identified by key (both type and value), so If you need private data
// just create a type that isn't exported from your package.
//
// That way, all items of that Type are accessible only inside your package
//
// Inspiration from Brad Fitizpatrick: http://groups.google.com/group/golang-nuts/msg/e2d679d303aa5d53
type Context struct {
	data map[interface{}]interface{}
	l sync.RWMutex
}

// Read the key from the context.
func (c *Context) Get(key interface{}) (v interface{}, has bool) {
	c.l.RLock()
	defer c.l.RUnlock()
	v, has = c.data[key]
	return
}

// Store the data on the context, if it already had it the old value is returned
func (c *Context) Set(key interface{}, val interface{}) (old interface{}, has bool) {
	c.l.Lock()
	defer c.l.Unlock()
	old, has = c.Get(key)
	c.data[key] = val
	return
}

// Used to define private keys
type privateKey *http.Request

var (
	// internal context used to store request context
	internalCtx = NewContext()
)

// Create a new empty context
func NewContext() *Context {
	ctx := new(Context)
	ctx.data = make(map[interface{}]interface{})
	return ctx
}

// Open the context
func OpenCtx(k *http.Request) *Context {
	pk := privateKey(k)
	if ctx, has := internalCtx.Get(pk); has {
		return ctx.(*Context)
	}
	return nil
}

// Remove the context associated with the request
func closeCtx(k *http.Request) {
	pk := privateKey(k)
	if _, has := internalCtx.Get(pk); has {
		// remove from the map
		delete(internalCtx.data, pk)
	}
}

// Bind a new context for the life-time of the request handling process.
//
// The user can call Open(req) as many times as he wants and only one instance will be created per request object
func BindContext(hndl http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		OpenCtx(req)
		defer closeCtx(req)
		hndl.ServeHTTP(w, req)
	})
}
