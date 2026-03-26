package web

import "net/http"

func (s *CryptoMonitorServer) showDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.templ.Render(w, s.prices, "dashboard.html"); err != nil {
		s.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *CryptoMonitorServer) refreshHandler(w http.ResponseWriter, r *http.Request) {
	s.prices = s.GetPrices()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
