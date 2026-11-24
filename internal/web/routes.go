package web

import (
	"encoding/json"
	"net/http"
)

// registerRoutes connects browser pages and JSON API endpoints.
// TODO: Add authentication checks to routes that read or update user data.
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "menu.tmpl", pageData{Title: "Space Adventure"})
	})
	mux.HandleFunc("GET /story", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "story.tmpl", pageData{Title: "Your Story"})
	})
	mux.HandleFunc("GET /ending", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "ending.tmpl", pageData{Title: "Story Ending"})
	})
	mux.HandleFunc("GET /achievements", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "achievements.tmpl", pageData{Title: "Achievements"})
	})

	mux.HandleFunc("GET /api/stories", notImplemented("Load story nodes and choices from Firebase."))
	mux.HandleFunc("GET /api/achievements", notImplemented("Load the signed-in user's achievements from Firebase."))
}

// notImplemented gives API routes a clear JSON response until their data logic exists.
func notImplemented(todo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotImplemented)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "This endpoint is not implemented yet.",
			"todo":  todo,
		})
	}
}
