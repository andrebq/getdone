package main

import (
	"flag"
	"github.com/andrebq/getdone/log"
	"github.com/andrebq/getdone/web"
	"net/http"
	"os"
)

var (
	start  *string = flag.String("s", "web", "What kind of app to start")
	port   *string = flag.String("port", ":8080", "Port to listen for http connections")
	prefix *string = flag.String("prefix", "", "Prefix of the web application")
	root   *string = flag.String("root", "./", "Root path to load contents. Must be the \"site\" folder.")
	help   *bool   = flag.Bool("h", false, "Show this help")
	l              = &log.Log{}
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	switch *start {
	case "web":
		runWeb()
	default:
		usage()
	}
}

func usage() {
	flag.Usage()
}

func runWeb() {
	l.Info("Preparing routes")
	l.Info("Root folder: %v", *root)
	rootHandler := web.Root(*root, *prefix)
	l.Info("Connecting to mongo server")
	err := web.InitMongo("localhost");
	if err != nil {
		l.Error("Error while connecting to mongo db. %v", err)
		os.Exit(1)
	}
	l.Info("Starting server @ %v", *port)
	err = http.ListenAndServe(*port, rootHandler)
	if err != nil {
		l.Error("Error while starting server. %v", err)
	} else {
		l.Info("Servidor atualizado com sucesso")
	}
}
