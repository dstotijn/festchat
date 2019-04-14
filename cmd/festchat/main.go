package main

import (
	"log"
	"net/http"
	"os"

	httphandler "github.com/dstotijn/festchat/handler/http"
	"github.com/dstotijn/festchat/sms/messagebird"
	"github.com/dstotijn/festchat/storage/memory"
)

func main() {

	mbAccessKey := os.Getenv("MESSAGEBIRD_ACCESS_KEY")
	if mbAccessKey == "" {
		log.Fatal("MESSAGEBIRD_ACCESS_KEY must be set.")
	}

	memStore := memory.NewStore()
	messagebird := messagebird.NewClient(mbAccessKey)

	handler := httphandler.NewHandler(messagebird, memStore)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
