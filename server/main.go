package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8888", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/main_go.js":
			w.Header().Set("Content-Type", "application/javascript")
			http.ServeFile(w, r, "./main_go.js")
		case "/main.wasm":
			w.Header().Set("Content-Type", "application/wasm")
			http.ServeFile(w, r, "../webapp/main.wasm")
		case "/":
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, indexHTML, "Hello, world!")
		default:
			http.FileServer(http.Dir("../webapp")).ServeHTTP(w, r)
		}
	}))
}

const indexHTML = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>%s</title>
	</head>
	<body>
		<script src="main_go.js"></script>
	</body>
</html>
`
