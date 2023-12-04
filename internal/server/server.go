package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"templates/internal/api"
	v1server "templates/swagger/v1"
)

type Server struct {
	router http.Server
	add    string
}

func (srv *Server) New() {
	router := mux.NewRouter()
	router.Handle("/v1", v1server.Handler(&api.V1{}))
	srv.router.Handler = router
	srv.add = ":9999"

}

func (srv *Server) Start() {
	go func() {
		if err := srv.router.ListenAndServe(); err != nil {
			fmt.Errorf("error start server: %s", err.Error())
		}
	}()
}
