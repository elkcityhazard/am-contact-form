package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/elkcityhazard/am-contact-form/internal/models"
	"github.com/justinas/nosurf"
)

func (m *Repository) HandleFormSubmission(w http.ResponseWriter, r *http.Request) {

	var payload = map[string]interface{}{}

	key := m.App.Router.GetField(r, 0)

	payload["key"] = key

	var msg models.Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		payload["error"] = err.Error()

		if err = json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	}

	if !nosurf.VerifyToken(nosurf.Token(r), msg.Token) {
		fmt.Println(nosurf.Token(r), msg.Token)
		fmt.Println("could not verify token")
	}

	msg.CreatedAt = time.Now()
	msg.UpdatedAt = time.Now()
	msg.Version = 1

	//validate

	var errors = map[string][]string{}

	if msg.Name == "" {
		errors["name"] = append(errors["name"], "name must not be empty")
	}

	if len(msg.Name) < 2 {
		errors["name"] = append(errors["name"], "name must not be a single character")
	}

	msg.Name = strings.ToLower(html.EscapeString(msg.Name)) // escape naughty hmtl

	_, err = mail.ParseAddress(msg.Email)

	if err != nil {
		errors["email"] = append(errors["email"], "you provided an invalid email address")
	}

	if msg.Email == "" {
		errors["email"] = append(errors["email"], "email cannot be blank")
	}

	msg.Email = strings.ToLower(html.EscapeString(msg.Email))

	if msg.MessageContent == "" {
		errors["message_content"] = append(errors["message_content"], "you didn't provide a message")
	}

	msg.MessageContent = html.EscapeString(msg.MessageContent)

	id, err := m.DB.InsertMessage(&msg)

	if err != nil {

		payload["error"] = err.Error()

		if err = json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	}

	payload["id"] = id
	payload["message"] = msg

	if err = json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
