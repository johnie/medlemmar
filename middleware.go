package main

import (
	"net/http"
)

func CORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Allow")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method != "OPTIONS" {
			h.ServeHTTP(w, r)
		}
	}
}
