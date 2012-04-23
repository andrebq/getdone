package web

import (
	"github.com/andrebq/getdone/log"
	"github.com/bmizerany/pat"
	"encoding/json"
	"net/http"
	"fmt"
)

var (
	rootFolder string = "./"
	l                 = &log.Log{}
)

// Root handler
func Root(root, prefix string) http.Handler {
	rootFolder = root
	mux := pat.New()
	mux.Get("/script/", http.HandlerFunc(LoadAsset))
	mux.Get("/style/", http.HandlerFunc(LoadAsset))
	mux.Get("/", http.HandlerFunc(ServeStatic))

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

// Serialize the data object into a JsonObject and write the contents to the reponse
//
// This call doesn't use the cuncked Transfer-Encoding since it will set the Content-Length header
// If the mimetype isn't specified (ie mimetype == "") application/json is used
func WriteJson(w http.ResponseWriter, data interface{}, mimetype string, okStatus int) (n int, err error) {
	tmp, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(tmp)))
	w.WriteHeader(okStatus)
	n, err = w.Write(tmp)
	return
}
