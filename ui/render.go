package ui

import (
	"embed"
	"html/template"
	"io"

	"cryptomonitor/internal/domain"
)

var (
	//go:embed "html"
	files embed.FS
)

func Render(w io.Writer, coins []domain.Price) error {
	templ, err := template.ParseFS(files, "html/templates/dashboard.html")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, coins); err != nil {
		return err
	}

	return nil
}
