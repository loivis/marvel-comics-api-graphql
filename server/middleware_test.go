package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	t.Run("OPTIONS", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodOptions, "/", nil)

		var called bool

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
		})

		cors(h).ServeHTTP(rr, req)

		if got, want := rr.Code, http.StatusOK; got != want {
			t.Errorf("got status code %d, want %d", got, want)
		}

		if called {
			t.Errorf("original handler was called")
		}

		if got, want := rr.Header().Get("Access-Control-Allow-Headers"), "*"; got != want {
			t.Errorf("got CORS header %q, want %q", got, want)
		}

		if got, want := rr.Header().Get("Access-Control-Allow-Origin"), "*"; got != want {
			t.Errorf("got CORS header %q, want %q", got, want)
		}
	})

	t.Run("GET", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		var called bool

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
		})

		cors(h).ServeHTTP(rr, req)

		if got, want := rr.Code, http.StatusOK; got != want {
			t.Errorf("got status code %d, want %d", got, want)
		}

		if !called {
			t.Errorf("original handler was NOT called")
		}

		if got, want := rr.Header().Get("Access-Control-Allow-Headers"), ""; got != want {
			t.Errorf("got CORS header %q, want %q", got, want)
		}

		if got, want := rr.Header().Get("Access-Control-Allow-Origin"), "*"; got != want {
			t.Errorf("got CORS header %q, want %q", got, want)
		}
	})
}
