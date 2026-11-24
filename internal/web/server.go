package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ScotlynTurner/space-adventure/internal/store"
)

type pageData struct {
	Title string
	Story *store.StoryNode
}

// NewServer creates the page routes and serves the frontend assets.
func NewServer() http.Handler {
	mux := http.NewServeMux()
	registerRoutes(mux)
	return mux
}

func renderPage(w http.ResponseWriter, page string, data pageData) {
	templates, err := template.ParseFiles(
		filepath.Join("internal", "view", "templates", "layout.tmpl"),
		filepath.Join("internal", "view", "templates", page),
	)
	if err != nil {
		log.Printf("load page template %q: %v", page, err)
		http.Error(w, "Page template unavailable", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(w, "layout.tmpl", data); err != nil {
		log.Printf("render page template %q: %v", page, err)
	}
}
