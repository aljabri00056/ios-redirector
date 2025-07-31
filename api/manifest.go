package handler

import (
	"fmt"
	"log"
	"net/http"
)

const manifestTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>items</key>
    <array>
        <dict>
            <key>assets</key>
            <array>
                <dict>
                    <key>kind</key>
                    <string>software-package</string>
                    <key>url</key>
                    <string>%s</string>
                </dict>
            </array>
            <key>metadata</key>
            <dict>
                <key>bundle-identifier</key>
                <string>%s</string>
                <key>bundle-version</key>
                <string>%s</string>
                <key>kind</key>
                <string>software</string>
                <key>title</key>
                <string>%s</string>
            </dict>
        </dict>
    </array>
</dict>
</plist>`

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[manifest] Request from %s with params: %s", r.RemoteAddr, r.URL.RawQuery)

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

	manifest := fmt.Sprintf(manifestTemplate, ipaURL, bundleID, bundleVersion, title)

	log.Printf("[manifest] Generated manifest for bundle_id: %s, title: %s, bundle_version: %s", bundleID, title, bundleVersion)

	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(manifest))
}