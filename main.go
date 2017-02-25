package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

type methodMux map[string]http.Handler

func (m *methodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := (*m)[r.Method]; ok {
		h.ServeHTTP(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	listen := os.Getenv("LISTEN")
	root := os.Getenv("ROOT")
	prefix := os.Getenv("PREFIX")

	files := http.StripPrefix(prefix, http.FileServer(http.Dir(root)))
	webdav := &webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.Dir(root),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("r=%v err=%v", r, err)
			}
		},
	}
	mux := methodMux(map[string]http.Handler{
		"GET":       files,
		"OPTIONS":   webdav,
		"PROPFIND":  webdav,
		"PROPPATCH": webdav,
		"MKCOL":     webdav,
		"COPY":      webdav,
		"MOVE":      webdav,
		"LOCK":      webdav,
		"UNLOCK":    webdav,
		"DELETE":    webdav,
		"PUT":       webdav,
	})

	if err := http.ListenAndServe(listen, &mux); err != nil {
		log.Fatal(err)
	}
}
