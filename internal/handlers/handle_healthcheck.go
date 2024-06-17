package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repo) HandleHealthcheck(w http.ResponseWriter, r *http.Request) {
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
