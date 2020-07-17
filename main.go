package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

const configLocation = "/etc/portfolio/"

var log *Logger

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
