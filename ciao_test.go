package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirection(t *testing.T) {
	type test struct {
		redirect string
		location string
		code     int
	}

	tests := []test{
		{
			redirect: "blog.creekorful.com",
			location: "https://blog.creekorful.dev",
			code:     308,
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Host = test.redirect

		c := &Config{Redirects: map[string]Redirect{
			test.redirect: {
				Location: test.location,
				Code:     test.code,
			},
		}}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(redirectHandler(c))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.code {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.code)
		}
		if location := rr.Header().Get("Location"); location != test.location {
			t.Errorf("handler returned wrong status location: got %v want %v", location, test.location)
		}
	}
}
