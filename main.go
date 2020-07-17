package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var log *Logger

func main() {
	log, f := NewLogger()

	router := httprouter.New()
	log.LogInfo("Creating new server")
	server, err := NewServer()
	if err != nil {
		log.LogWarningf("Could not create server from config, creating default. Error %v", err)
		server, err = CreateConfig()
		if err != nil {
			log.LogErrorf("Could not create default server instance. Error %v", err)
			Error(router, err)
		}
	}
	server.Router = router
	router.ServeFiles("/static/*filepath", http.Dir("/etc/portfolio/assets"))
	router.GET("/", server.Index)
	log.LogInfo("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
	//	f.Close()
}

func (srv *Server) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("/etc/portfolio/assets/index.html")
	t.Execute(w, srv)
}

func Error(router *httprouter.Router, err error) {
	router.GET("/error", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, err)
	})

}
