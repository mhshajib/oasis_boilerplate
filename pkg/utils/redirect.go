package utils

import (
	"net/http"
	"net/url"
)

func Redirect(w http.ResponseWriter, r *http.Request, redirectParameters map[string]string, parsedURL *url.URL) {
	parameters := url.Values{}
	for key, value := range redirectParameters {
		parameters.Add(key, value)
	}
	parsedURL.RawQuery = parameters.Encode()
	redirectURL := parsedURL.String()
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}
