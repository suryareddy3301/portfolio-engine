package main

import (
	"fmt"
	l "log"
	"net/http"
	"os"
	"runtime"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var configLocation, logLocation string

//const configLocation = "/etc/portfolio/"

var log *Logger

func init() {
	if runtime.GOOS == "windows" {
		err := os.MkdirAll("C:/portfolio/config/", os.ModePerm)
		if err != nil {
			l.Fatal(err)
		}
		err = os.MkdirAll("C:/portfolio/logs/", os.ModePerm)
		if err != nil {
			l.Fatal(err)
		}
		//configLocation, logLocation = "C:/portfolio/config/", "C:/portfolio/logs/"
		configLocation, logLocation = "", "C:/portfolio/logs/"
	} else {
		configLocation, logLocation = "/etc/portfolio/", "/var/log/"
	}
}

func main() {
	log, file := NewLogger()
	defer file.Close()
	router := httprouter.New()
	log.LogInfo("Creating new server")
	server, err := NewServer()
	if err != nil {
		log.LogWarningf("Could not create a new server with the provide config. Error %v", err)
		log.LogInfo("Creating default config")
		server, err = CreateConfig()
		if err != nil {
			log.LogErrorf("could not create server with the default config. Error %v", err)
			Error(router, err)
		}
	}
	server.Router = router
	router.ServeFiles("/static/*filepath", http.Dir(configLocation+"assets"))
	router.GET("/", server.Index)

	log.LogFatal(http.ListenAndServe(":8080", router))
	//	defer file.Close()

}

func (srv *Server) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles(configLocation + "assets/index.html")
	t.Execute(w, srv)
}

func Error(router *httprouter.Router, err error) {
	router.GET("/error", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, err)
	})

}
