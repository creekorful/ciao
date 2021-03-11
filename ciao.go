package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

type Context struct {
	Redirects map[string]string `json:"redirects"`
}

func main() {
	configFlag := flag.String("config", "ciao.json", "path to the config file")

	flag.Parse()

	f, err := os.Open(*configFlag)
	if err != nil {
		panic(err)
	}

	var ctx Context
	if err := json.NewDecoder(f).Decode(&ctx); err != nil {
		panic(err)
	}
	_ = f.Close()

	log.Printf("Successfully loaded %d redirects", len(ctx.Redirects))

	http.HandleFunc("/", redirectHandler(&ctx))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func redirectHandler(c *Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if target, exist := c.Redirects[r.Host]; exist {
			log.Printf("%s - Redirecting %s -> %s", r.RemoteAddr, r.Host, target)
			w.Header().Add("Location", target)
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			log.Printf("%s - No redirect found for: %s", r.RemoteAddr, r.Host)
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
