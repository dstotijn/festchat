package http

import (
	"net/http"

	"github.com/dstotijn/festchat/chat"
	"github.com/dstotijn/festchat/sms"
)

// Handler represents an HTTP handler.
type Handler struct {
	mux     *http.ServeMux
	chatSvc *chat.Service
}

// NewHandler returns a new handler.
func NewHandler(smsSender sms.Sender, repo chat.Repository) *Handler {
	handler := &Handler{
		mux:     http.NewServeMux(),
		chatSvc: chat.NewService(smsSender, repo),
	}

	handler.mux.HandleFunc("/svc/received-sms", handler.receivedSMS)

	return handler
}

// ServeHTTP implements http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
