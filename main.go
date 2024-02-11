package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello From OffGridChat API\n")
}
func getAPIHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello\n")
}

func main() {
    http.HandleFunc("/", getRoot)
	http.HandleFunc("/api", getAPIHello)

	err := http.ListenAndServe(":5000", nil)

    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
