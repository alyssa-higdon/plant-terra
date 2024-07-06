package server

import (
	"fmt"
	"net/http"
	"text/template"
)

type ServerState struct {
}

var indexTpl = template.Must(template.ParseFiles("index.html"))

func (s *ServerState) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := indexTpl.Execute(w, nil); err != nil {
		fmt.Printf("Failed to execute index.html %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
