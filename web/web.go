package web

import (
	"github.com/andrebq/getdone/log"
	"github.com/bmizerany/pat"
	"net/http"
)

var (
	rootFolder string = "./"
	l                 = &log.Log{}
)

// Root handler
func Root(root, prefix string) http.Handler {
	rootFolder = root
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(ServeStatic))
	mux.Get("/script", http.HandlerFunc(LoadAsset))
	mux.Get("/style", http.HandlerFunc(LoadAsset))

	rootHndl := http.StripPrefix(prefix, mux)
	return Log(rootHndl)
}

// Log all requests
func Log(hndl http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		l.Info("%v %v", req.Method, req.URL)
		hndl.ServeHTTP(w, req)
	})
}
