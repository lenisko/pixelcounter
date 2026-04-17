package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed static
var static embed.FS

func main() {
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	port := flag.Int("port", 8373, "port to listen on")
	flag.Parse()

	http.Handle("/", http.FileServer(http.FS(sub)))

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Pixel Counter running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
