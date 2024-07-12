package main

import (
	"fmt"
	"net/http"
)

type serve int

func (m serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("lucas-Key", "This key is from Lucas")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code can run from this func</h1>")
}

func main() {
	var d serve
	http.ListenAndServe(":8080", d)
}
