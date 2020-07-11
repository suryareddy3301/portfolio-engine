package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	server := NewServer()
	server.Router = router
	router.ServeFiles("/static/*filepath", http.Dir("/etc/portfolio/assets"))
	router.GET("/", server.Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (srv *Server) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//srv.Router.ServeFiles("/etc/portfolio/assets/*filepath", http.Dir("/etc/portfolio/assets"))
	t, _ := template.ParseFiles("/etc/portfolio/assets/index.html")
	t.Execute(w, srv)
}
