package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elkcityhazard/am-contact-form/internal/mailer"
	"github.com/elkcityhazard/am-contact-form/internal/models"
	"github.com/elkcityhazard/am-contact-form/pkg/utils"
	"github.com/justinas/nosurf"
)

func (m *Repository) HandleFormSubmission(w http.ResponseWriter, r *http.Request) {
	payload := map[string]interface{}{}

	key := m.App.Router.GetField(r, 0)

	token, err := r.Cookie("csrf_token")
	if err != nil {
		payload["error"] = err.Error()

		if err = json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	}

	if !nosurf.VerifyToken(nosurf.Token(r), token.Value) {
		err = errors.New("could not validate token")
		if err != nil {
			payload["error"] = err.Error()

			if err = json.NewEncoder(w).Encode(payload); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return

		}

	}

	transaction, err := r.Cookie("transaction")
	if err != nil {
		payload["error"] = err.Error()

		if err = json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	}

	transactionVals := strings.SplitN(transaction.Value, "|", 3)

	isValidToken := utils.NewUtil().VerifyHmacToken(transaction.Value, "|")

	if !isValidToken {
		payload["error"] = errors.New("invalid token").Error()

		if err = json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	payload["key"] = key
	payload["transaction"] = transactionVals

	payload["sent_token"] = token.Value
	payload["actual_token"] = nosurf.Token(r)
	payload["ip_address"] = ReadUserIP(r)
	var msg models.Message

	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		payload["error"] = err.Error()

		if err = json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	}

	msg.Token = token.Value

	csrf_token, err := r.Cookie("csrf_token")
	if err != nil {
		fmt.Println(err.Error())
	}

	if !nosurf.VerifyToken(nosurf.Token(r), csrf_token.Value) {
		err := errors.New("bad request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	msg.CreatedAt = time.Now()
	msg.UpdatedAt = time.Now()
	msg.Version = 1

	// validate

	errors := map[string][]string{}

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
	msg.IP = ReadUserIP(r)

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

	payload["htmlBody"] = msg
	payload["plainBody"] = msg
	payload["SiteName"] = "andrew-mccall.com"

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		port = 0
	}
	mailer := mailer.New(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_USERNAME"))

	err = mailer.SendEmail("andrew@andrew-mccall.com", "submission.tmpl", payload)
	if err != nil {

		payload["error"] = err.Error()

		if err = json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	}

	if err = json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
