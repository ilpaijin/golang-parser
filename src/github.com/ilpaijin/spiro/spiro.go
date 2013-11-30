package main

import (
	"github.com/ajstarks/svgo"
	// "io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/spiro", http.HandlerFunc(SpiroHandler))
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func SpiroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	Circle(w)
}

func Circle(w http.ResponseWriter) {
	canvas := svg.New(w)
	canvas.Start(500, 500)
	canvas.Circle(100, 50, 40, "fill:red; stroke:black; stroke-width:2")
	canvas.End()
}
