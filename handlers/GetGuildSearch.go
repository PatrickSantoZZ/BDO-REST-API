package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bdo-rest-api/scrapers"
)

func GetGuildSearch(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	regionParams, ok1 := r.URL.Query()["region"]
	pageParams, ok2 := r.URL.Query()["page"]
	queryParams, ok3 := r.URL.Query()["query"]

	if !ok1 || !validateRegion(regionParams[0]) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page := 1

	if ok2 {
		page, _ = strconv.Atoi(pageParams[0])
	}

	var query string

	if ok3 {
		query = queryParams[0]
	}

	if data, status := scrapers.ScrapeGuildSearch(regionParams[0], query, int32(page)); status == http.StatusOK {
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(status)
	}
}
