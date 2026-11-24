package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type pageData struct {
	Title string
}

// NewServer builds the HTTP routes and serves the app's static assets.
func NewServer() http.Handler {
	mux := http.NewServeMux()
	registerRoutes(mux)

	mux.Handle("GET /styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("internal/view/styles"))))
	mux.Handle("GET /scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("internal/view/scripts"))))
	mux.Handle("GET /images/", http.StripPrefix("/images/", http.FileServer(http.Dir("static/images"))))
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
