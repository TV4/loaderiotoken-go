package loaderiotoken

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		mux := http.NewServeMux()

		Register(mux, "foo-token")

		ts := httptest.NewServer(mux)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/loaderio-foo-token.txt")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Errorf("resp.StatusCode = %d, want %d", got, want)
		}

		if got, want := resp.Header.Get("Cache-Control"), "public, max-age=10"; got != want {
			t.Errorf("Cache-Control header = %q, want %q", got, want)
		}

		gotBody, _ := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := gotBody, []byte("loaderio-foo-token\n"); !bytes.Equal(got, want) {
			t.Errorf("body = %q, want %q", got, want)
		}
	})

	t.Run("EmptyToken", func(t *testing.T) {
		mux := http.NewServeMux()

		emptyToken := ""

		Register(mux, emptyToken)

		ts := httptest.NewServer(mux)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/loaderio-.txt")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := resp.StatusCode, http.StatusNotFound; got != want {
			t.Errorf("resp.StatusCode = %d, want %d", got, want)
		}
	})
}
