package handlers

import "net/http"

func validateRegion(r string) bool {
	return r == "EU" || r == "NA"
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
