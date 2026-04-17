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

	host := flag.String("host", "localhost", "host to bind to")
	port := flag.Int("port", 8373, "port to listen on")
	flag.Parse()

	http.Handle("/", http.FileServer(http.FS(sub)))

	addr := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Pixel Counter running at http://%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
