package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":8888", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/main.wasm" {
			log.Printf("serving wasm %q", r.URL.Path)
			w.Header().Set("Content-Type", "application/wasm")
			buf, err := ioutil.ReadFile("../webapp/main.wasm")
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(buf)
			return
		}
		http.FileServer(http.Dir("../webapp")).ServeHTTP(w, r)
	}))
}
