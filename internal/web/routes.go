package web

import (
	"net/http"

	"github.com/ScotlynTurner/space-adventure/internal/store"
)

// registerRoutes connects the app's pages and static assets.
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "menu.tmpl", pageData{Title: "Space Adventure"})
	})
	mux.HandleFunc("GET /story", func(w http.ResponseWriter, r *http.Request) {
		// Temporary sample data lets the story page work before Firestore is connected.
		story := store.StoryNode{
			Description: "Your ship's alarm wakes you from cryosleep. Two warning lights flash on the console.",
			Choices: []store.Choice{
				{Label: "Check the cockpit", NextNodeID: "cockpit"},
				{Label: "Investigate the cargo bay", NextNodeID: "cargo-bay"},
			},
		}
		renderPage(w, "story.tmpl", pageData{Title: "Your Story", Story: &story})
	})
	mux.HandleFunc("GET /ending", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "ending.tmpl", pageData{Title: "Story Ending"})
	})
	mux.HandleFunc("GET /achievements", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "achievements.tmpl", pageData{Title: "Achievements"})
	})

	mux.Handle("GET /styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("internal/view/styles"))))
	mux.Handle("GET /scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("internal/view/scripts"))))
	mux.Handle("GET /images/", http.StripPrefix("/images/", http.FileServer(http.Dir("static/images"))))
}
