package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleGetCSRFToken(w http.ResponseWriter, r *http.Request) {

	var payload = map[string]interface{}{}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
