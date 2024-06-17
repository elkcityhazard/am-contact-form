package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repo) HandleFormSubmission(w http.ResponseWriter, r *http.Request) {
	key := activeRouter.GetField(r, 0)

	http.SetCookie(w, &http.Cookie{
		Name:     "id",
		Value:    "1",
		Path:     "/",
		MaxAge:   60,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	err := json.NewEncoder(w).Encode(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
