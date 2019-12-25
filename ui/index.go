package ui

import (
	"net/http"
	"regexp"
)

// IndexHandler ..
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	apiPath := "^/api/([a-zA-Z0-9]+)$"
	matched, err := regexp.Match(apiPath, []byte(r.URL.Path))

	if !matched {
		http.ServeFile(w, r, "static/index.html")
	} else {
		http.NotFound(w, r)
	}

	if err != nil {
		panic(err)
	}
}
