package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/templates/internal/api"
	v1server "github.com/templates/swagger/v1"
	"net/http"
)

type Server struct {
	router *http.Server
	add    string
}

func (srv *Server) New() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(v1server.HandlerWithOptions(&api.V1{}, v1server.GorillaServerOptions{BaseURL: "/v1"}))
	srv.add = "127.0.0.1:8080"
	srv.router = &http.Server{
		Addr:    srv.add,
		Handler: router,
	}
}

func (srv *Server) Start() {
	go func() {
		fmt.Println("Started server at: ", srv.add)
		if err := srv.router.ListenAndServe(); err != nil {
			fmt.Errorf("error start server: %s", err.Error())
		}
	}()
}
