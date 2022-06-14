package middleware

import "net/http"

func (api) AllowOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Acces-Control-Allow-Origin", "http.Request.Header.Get(Origin)")
	w.Header().Set("Acces-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}
