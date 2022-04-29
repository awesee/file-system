package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	DIR  string
	PORT uint
)

func init() {
	flag.StringVar(&DIR, "dir", ".", "directory to serve")
	flag.UintVar(&PORT, "port", 8090, "port to serve on")
	flag.Parse()
}

func main() {
	port := fmt.Sprintf(":%d", PORT)
	fmt.Println("http://127.0.0.1" + port)
	fs := http.FileServer(http.Dir(DIR))
	log.Fatalln(http.ListenAndServe(port, fs))
}
