package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Config is the application configuration
type Config struct {
	Redirects     map[string]Redirect `json:"redirects"`
	UseXForwarded bool                `json:"use_x_forwarded"`
}

// Redirect represent a redirection rule
type Redirect struct {
	Location string `json:"location"`
	Code     int    `json:"code"`
}

func (r *Config) findRedirect(host string, u *url.URL) (Redirect, bool) {
	// first of all check if there's an exact rule available (host + path)
	if val, exist := r.Redirects[host+u.Path]; exist {
		return val, true
	}

	// otherwise fallback to host rule
	val, exist := r.Redirects[host]
	return val, exist
}

func main() {
	configFlag := flag.String("config", "direktion.json", "path to the config file")

	flag.Parse()

	f, err := os.Open(*configFlag)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}
	_ = f.Close()

	log.Printf("Successfully loaded %d redirections", len(config.Redirects))

	http.HandleFunc("/", redirectHandler(&config))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func redirectHandler(c *Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if redirect, exist := c.findRedirect(r.Host, r.URL); exist {
			code := http.StatusTemporaryRedirect
			if redirect.Code != 0 {
				code = redirect.Code
			}

			remoteIP := r.RemoteAddr
			if c.UseXForwarded {
				remoteIP = getRealIP(r)
			}

			log.Printf("%s - [%d] Redirecting %s%s -> %s", remoteIP, code, r.Host, r.URL.Path, redirect.Location)

			w.Header().Add("Location", redirect.Location)
			w.WriteHeader(code)
		} else {
			log.Printf("%s - No redirect found for: %s", r.RemoteAddr, r.Host)
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getRealIP(r *http.Request) string {
	address := r.Header.Get("X-Real-Ip")
	if address == "" {
		address = r.Header.Get("X-Forwarded-For")
	}
	if address == "" {
		address = r.RemoteAddr
	}
	return address
}
