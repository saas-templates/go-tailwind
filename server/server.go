package server

import (
	"context"
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const gracePeriod = 5 * time.Second

var (
	//go:embed static/**
	staticFS embed.FS

	//go:embed templates/**
	templateFS embed.FS
)

// Serve starts the server on given `addr`.
func Serve(ctx context.Context, addr string) error {
	tpl, err := template.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		panic(err)
	}

	app := &App{
		tpl: tpl,
	}

	r := chi.NewRouter()
	r.Mount("/static", http.FileServer(http.FS(staticFS)))

	// Page Routes
	r.Get("/", app.indexPage)

	return graceServe(ctx, addr, r)
}

type App struct {
	tpl *template.Template
}

func (app *App) indexPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, "index.html", nil)
}

func (app *App) render(w http.ResponseWriter, name string, data any) {
	err := app.tpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("[WARN] failed to render template '%s': %v", name, err)
	}
}
