package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	serverDescription := map[string]string{
		"version": "1.0",
		"name":    "Hello World Server",
		"port":    ":3000",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(serverDescription)
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	fmt.Printf("Starting %s on port %s", serverDescription["name"], serverDescription["port"])
	http.ListenAndServe(serverDescription["port"], n)
}
