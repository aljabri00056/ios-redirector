package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func TrollstoreHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[trollstore] Request from %s with params: %s", r.RemoteAddr, r.URL.RawQuery)

	ipaURL := r.URL.Query().Get("ipa_url")

	if ipaURL == "" {
		http.Error(w, "Missing required parameter: ipa_url", http.StatusBadRequest)
		return
	}

	magnifierURL := fmt.Sprintf("apple-magnifier://install?url=%s", url.QueryEscape(ipaURL))

	log.Printf("[trollstore] Redirecting to: %s", magnifierURL)

	w.Header().Set("Location", magnifierURL)
	w.WriteHeader(http.StatusFound)
}