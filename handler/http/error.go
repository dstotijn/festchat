package http

import "net/http"

func badRequest(w http.ResponseWriter, error string) {
	code := http.StatusBadRequest
	http.Error(w, error, code)
}
