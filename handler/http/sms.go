package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dstotijn/festchat/chat"
)

type receivedMessage struct {
	ID         string    `json:"id"`
	Originator string    `json:"originator"`
	Recipient  string    `json:"recipient"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"createdDatetime"`
}

func (h Handler) receivedSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
		return
	}

	var msg receivedMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		badRequest(w, err.Error())
		return
	}

	go func() {
		err := h.chatSvc.RelayMessage(chat.ReceivedMessage{
			Originator: msg.Originator,
			Recipient:  msg.Recipient,
			Payload:    msg.Body,
		})
		if err != nil {
			log.Printf("Error: Could not relay received SMS: %v", err)
		}
	}()

	fmt.Fprintln(w, "OK")
}
