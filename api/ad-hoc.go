package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ad_hoc] Request from %s with params: %s", r.RemoteAddr, r.URL.RawQuery)

	ipaURL := r.URL.Query().Get("ipa_url")
	bundleID := r.URL.Query().Get("bundle_id")
	bundleVersion := r.URL.Query().Get("bundle_version")
	title := r.URL.Query().Get("title")

	if ipaURL == "" {
		http.Error(w, "Missing required parameter: ipa_url", http.StatusBadRequest)
		return
	}
	if bundleID == "" {
		http.Error(w, "Missing required parameter: bundle_id", http.StatusBadRequest)
		return
	}
	if bundleVersion == "" {
		http.Error(w, "Missing required parameter: bundle_version", http.StatusBadRequest)
		return
	}
	if title == "" {
		http.Error(w, "Missing required parameter: title", http.StatusBadRequest)
		return
	}

	host := r.Host
	scheme := "https"
	if r.Header.Get("X-Forwarded-Proto") != "" {
		scheme = r.Header.Get("X-Forwarded-Proto")
	}

	manifestURL := fmt.Sprintf("%s://%s/api/manifest?ipa_url=%s&bundle_id=%s&bundle_version=%s&title=%s",
		scheme, host,
		url.QueryEscape(ipaURL),
		url.QueryEscape(bundleID),
		url.QueryEscape(bundleVersion),
		url.QueryEscape(title))

	itmsURL := fmt.Sprintf("itms-services://?action=download-manifest&url=%s",
		url.QueryEscape(manifestURL))

	log.Printf("[ad_hoc] Redirecting to: %s", itmsURL)

	w.Header().Set("Location", itmsURL)
	w.WriteHeader(http.StatusFound)
}