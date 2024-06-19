package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/nosurf"
)

func (m *Repository) HandleHealthcheck(w http.ResponseWriter, r *http.Request) {

	m.App.SessionManager.Put(r.Context(), "id", 1)

	type Healthcheck struct {
		Status  int    `json:"status"`
		Version string `json:"version"`
		Msg     string `json:"message"`
		Token   string `json:"token"`
	}

	var s Healthcheck

	s.Status = 200
	s.Version = ""
	s.Msg = "All Systems Checkout"
	s.Token = nosurf.Token(r)

	if err := json.NewEncoder(w).Encode(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
