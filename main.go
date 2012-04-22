package main

import (
	"flag"
	"net/http"
	"github.com/andrebq/getdone/web"
	"github.com/andrebq/getdone/log"
)

var (
	start *string = flag.String("s", "web", "What kind of app to start")
	port *string = flag.String("port", ":8080", "Port to listen for http connections")
	prefix *string = flag.String("prefix", "", "Prefix of the web application")
	root *string = flag.String("root", "./", "Root path to load contents. Must be the \"site\" folder.")
	help *bool = flag.Bool("h", false, "Show this help")
	l = &log.Log{}
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	switch(*start) {
		case "web": runWeb()
		default: usage()
	}
}

func usage() {
	flag.Usage()
}

func runWeb() {
	l.Info("Preparing routes")
	root := web.Root(*root, *prefix)
	l.Info("Starting server")
	err := http.ListenAndServe(*port, root)
	if err != nil {
		l.Error("Error while starting server. %v", err)
	} else {
		l.Info("Servidor atualizado com sucesso")
	}
}
