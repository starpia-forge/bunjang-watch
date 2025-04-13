package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"net/http"
	"path/filepath"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 정적 파일
	// TODO - 수정 필요
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// 페이지 라우팅
	r.Get("/", render("index.html"))

	// API 라우팅
	r.Route("/api", func(r chi.Router) {
		r.Post("/watcher", nil)
		r.Post("/hook", nil)
	})

	return r
}

func render(templateName string) http.HandlerFunc {
	// TODO - 수정 필요
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			filepath.Join("templates", "layout.html"),
			filepath.Join("templates", templateName),
		))
		tmpl.ExecuteTemplate(w, "layout", nil)
	}
}

func renderPartial(templateName string) http.HandlerFunc {
	// TODO - 수정 필요
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			filepath.Join("templates", templateName),
		))

		// HTMX 헤더 처리 가능 (선택)
		if r.Header.Get("HX-Request") == "true" {
			w.Header().Set("HX-Trigger", "partial-loaded")
		}

		tmpl.Execute(w, nil)
	}
}
