package web

import "net/http"

func (s *CryptoMonitorServer) showDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.templ.Render(w, s.prices, "dashboard.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *CryptoMonitorServer) refreshHandler(w http.ResponseWriter, r *http.Request) {
	s.RefreshPrices()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
