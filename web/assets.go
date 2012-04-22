package web

import (
	"net/http"
	"path"
	"path/filepath"
)

// Load the assets from the assets folder compiling if necessary, there are three kinds of assets that this handler can serve
//
// 1. CoffeeScript/JavaScript files. CoffeeScript files are compiled on-demand (and cached)
// 2. Less/CSS files. Less files are compiled on-demand (and cached)
// 3. Image files. Images are just serverd and don't get any pre-processing
func LoadAsset(w http.ResponseWriter, req *http.Request) {
	resPath := req.URL.Path
	switch path.Ext(resPath) {
	case ".js":
		serveJs(w, req)
	case ".coffee":
		serveCoffee(w, req)
	case ".css":
		serveCss(w, req)
	case ".less":
		serveLess(w, req)
	default:
		serveAsset(w, req)
	}
}

// Just serve static files from the site folder.
func ServeStatic(w http.ResponseWriter, req *http.Request) {
	resPath := ""
	if req.URL.Path == "/" {
		resPath = fixStatic("./index.html")
	} else {
		resPath = fixStatic(req.URL.Path)
	}

	resPath, err := filepath.Abs(filepath.FromSlash(resPath))
	if err != nil {
		l.Error("Error while fetching the Abs path for %v. %v", req.URL.Path, err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	l.Info("%v mapped to %v", req.URL.Path, resPath)
	http.ServeFile(w, req, filepath.FromSlash(resPath))
}

// Serve static files without any processing
func serveAsset(w http.ResponseWriter, req *http.Request) {
	assetPath := fixAsset(req.URL.Path)
	http.ServeFile(w, req, filepath.FromSlash(assetPath))
}

// just invoke the serveAsset
func serveJs(w http.ResponseWriter, req *http.Request) {
	serveAsset(w, req)
}

func serveCoffee(w http.ResponseWriter, req *http.Request) {
	panic("not done")
}

func serveLess(w http.ResponseWriter, req *http.Request) {
	panic("not done")
}

// just invoke the serveAsset
func serveCss(w http.ResponseWriter, req *http.Request) {
	serveAsset(w, req)
}

func fixAsset(file string) string {
	return path.Join(fixStatic(""), "assets", path.Clean(file))
}

func fixStatic(file string) string {
	if file == "" {
		return path.Join(rootFolder)
	}
	return path.Join(rootFolder, path.Clean(file))
}
