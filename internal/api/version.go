package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	defaultVersion = "UNKNOWN"
	defaultDate    = "UNKNOWN"
)

type Info struct {
	Version string `json:"version"`
	Date    string `json:"date"`
}

func (srv *V1) GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetVersion")
	i := Info{
		Version: defaultVersion,
		Date:    defaultDate,
	}
	res, err := json.Marshal(&i)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("error marshalling data: %v", err)))
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
