package ui

import (
	"regexp"
	"net/http"
)

var apiPath = "^/api/([a-zA-Z0-9]+)$"

// IndexHandler ..
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	matched, err := regexp.Match(apiPath, []byte(r.URL.Path))
	if !matched {
		http.ServeFile(w, r, "static/index.html")
	}
	if err != nil {
		panic(err)
	}
}
