package web

import "net/http"

func (s *CryptoMonitorServer) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /{$}", s.showDashboardHandler)
	router.HandleFunc("POST /refresh", s.refreshHandler)
	return router
}
