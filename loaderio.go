// Package loaderiotoken sets up a verification endpoint for loader.io
package loaderiotoken

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Register registers an endpoint /loaderio-<token>.txt, used by loader.io for
// verification. No endpoint will be registered if the token is blank.
func Register(mux *http.ServeMux, token string) {
	if token == "" {
		return
	}
	mux.HandleFunc(
		fmt.Sprintf("/loaderio-%s.txt", url.PathEscape(token)),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "public, max-age=10")
			io.WriteString(w, fmt.Sprintf("loaderio-%s\n", token))
		},
	)
}
