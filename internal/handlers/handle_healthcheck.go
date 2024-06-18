package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repository) HandleHealthcheck(w http.ResponseWriter, r *http.Request) {

	m.App.SessionManager.Put(r.Context(), "id", 1)

	type Healthcheck struct {
		Status  int    `json:"status"`
		Version string `json:"version"`
		Msg     string `json:"message"`
	}

	var s Healthcheck

	s.Status = 200
	s.Version = ""
	s.Msg = "All Systems Checkout"

	if err := json.NewEncoder(w).Encode(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
