package main

import (
	"log"
	"net/http"

	"github.com/ScotlynTurner/space-adventure/internal/web"
)

func main() {
	config := loadConfig()
	handler := web.NewServer()

	log.Printf("Space Adventure listening on %s", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, handler))
}
