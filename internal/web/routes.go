package web

import (
	"net/http"

	"cryptomonitor/ui"
)

func (s *CryptoMonitorServer) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /static/", http.FileServerFS(ui.Files))

	router.HandleFunc("GET /{$}", s.showDashboardHandler)
	router.HandleFunc("POST /refresh", s.refreshHandler)

	return router
}
