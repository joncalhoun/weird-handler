package main

import (
	"fmt"
	"io"
	"os"

	"github.com/joncalhoun/weird-handler/http"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	http.HandleFunc("/", DemoHandler)
	return http.ListenAndServe(":3000", nil)
}

func DemoHandler(r *http.Request) http.Responder {
	var errCond bool = true
	if errCond {
		return http.Error("oh no", 500)
	}
	return http.OkResponder()
}
