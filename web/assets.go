package web

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
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

	resPath, err := fixAbs(w, req, resPath)
	// can return since fixAbs already set the headers in the response
	if err != nil {
		return
	}

	http.ServeFile(w, req, filepath.FromSlash(resPath))
}

// Convert the path to the system-depended absolut version of the path.
//
// If an error happens and the conversion can't be completed, set  404 status code. The mapping is logged for futher use
func fixAbs(w http.ResponseWriter, req *http.Request, path string) (resPath string, err error) {
	resPath, err = filepath.Abs(filepath.FromSlash(path))
	if err != nil {
		l.Error("Error while fetching the Abs path for %v. %v", req.URL.Path, err)
		http.Error(w, "", http.StatusNotFound)
		return
	}
	// if the don't have a .html sufix
	// then check if the directory or file exists
	// if not, check if the same file (with the .html prefix) exists
	// if true, return the .html path instead of a 404 error.
	if !strings.HasSuffix(resPath, ".html") {
		_, err := os.Open(resPath)
		if os.IsNotExist(err) {
			resPath = resPath + ".html"
		}
	}

	l.Printf("STATIC_MAP ", "%v mapped to %v", req.URL.Path, resPath)
	return
}

// Serve static files without any processing
func serveAsset(w http.ResponseWriter, req *http.Request) {
	assetPath, err := fixAbs(w, req, fixAsset(req.URL.Path))
	if err != nil {
		return
	}
	http.ServeFile(w, req, assetPath)
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
