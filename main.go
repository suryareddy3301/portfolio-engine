package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	server, err := NewServer()
	if err != nil {
		server, err = CreateConfig()
		if err != nil {
			Error(router, err)
		}
	}
	server.Router = router
	router.ServeFiles("/static/*filepath", http.Dir("/etc/portfolio/assets"))
	router.GET("/", server.Index)
	log.Fatal(http.ListenAndServe(":8080", router))
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
