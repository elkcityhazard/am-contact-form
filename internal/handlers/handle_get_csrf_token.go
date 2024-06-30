package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/elkcityhazard/am-contact-form/pkg/utils"
)

func HandleGetCSRFToken(w http.ResponseWriter, r *http.Request) {

	var payload = map[string]interface{}{}

	var u = utils.NewUtil()

	randomString := u.GenerateRandomString(24)

	token, err := u.CreateHmacToken(randomString)

	if err != nil {
		go panic(err)
	}

	payload["token"] = fmt.Sprintf("%s|%s|%d", token, randomString, time.Now().Add(time.Hour).Unix())

	http.SetCookie(w, &http.Cookie{
		Name:   "transaction",
		Value:  payload["token"].(string),
		Path:   "",
		Secure: Repo.App.IsProduction,
		MaxAge: 3600,
	})

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
