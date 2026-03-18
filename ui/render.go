package ui

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"

	"cryptomonitor/internal/domain"
)

var (
	//go:embed "html"
	files embed.FS
)

type Renderer struct {
	cache map[string]*template.Template
}

func NewRenderer() (*Renderer, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(files, "html/templates/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"html/base.html",
			page,
		}

		templ, err := template.ParseFS(files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = templ
	}

	return &Renderer{cache: cache}, nil
}

func (r *Renderer) Render(w io.Writer, coins []domain.Price, page string) error {
	templ, ok := r.cache[page]
	if !ok {
		return fmt.Errorf("El template %s no existe", page)
	}

	if err := templ.ExecuteTemplate(w, "base", coins); err != nil {
		return err
	}

	return nil
}
