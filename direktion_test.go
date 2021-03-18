package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirection(t *testing.T) {
	type test struct {
		url              string
		expectedLocation string
		expectedCode     int

		redirects map[string]Redirect
	}

	tests := []test{
		{
			url:              "https://blog.creekorful.com",
			expectedLocation: "https://blog.creekorful.dev",
			expectedCode:     307,
			redirects: map[string]Redirect{
				"blog.creekorful.com": {
					Location: "https://blog.creekorful.dev",
					Code:     307,
				},
			},
		},
		{
			url:              "https://blog.creekorful.com",
			expectedLocation: "",
			expectedCode:     404,
			redirects:        map[string]Redirect{},
		},
		{
			url:              "https://blog.creekorful.com/something-cool",
			expectedLocation: "https://something-cool.creekorful.com/blog",
			expectedCode:     307,
			redirects: map[string]Redirect{
				"blog.creekorful.com/something-cool": {
					Location: "https://something-cool.creekorful.com/blog",
					Code:     307,
				},
				"blog.creekorful.com": {
					Location: "https://blog.creekorful.dev",
					Code:     308,
				},
			},
		},
		{
			url:              "https://creekorful.me/2019/01/12/terminews",
			expectedLocation: "https://blog.creekorful.com/2019/01/12/terminews",
			expectedCode:     308,
			redirects: map[string]Redirect{
				"creekorful.me": {
					Location: "https://blog.creekorful.com/$request_uri",
					Code:     308,
				},
			},
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		c := &Config{Redirects: test.redirects}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(redirectHandler(c))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedCode)
		}
		if location := rr.Header().Get("Location"); location != test.expectedLocation {
			t.Errorf("handler returned wrong status location: got %v want %v", location, test.expectedLocation)
		}
	}
}
