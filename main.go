package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/text/encoding/unicode"
)

func main() {

	// hello is a hello world endpoint. It accepts a POST body of your name encoded
	// in utf-16, and returns "hello, <name>" in utf-8 encoding.
	// Java represents strings internally as utf-16, and Go represents them internally
	// as utf-8. This endpoint might help you make new Java and Go friends!
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "bad request")

			return
		}

		// convert the incoming utf-16 string to utf-8 -- neat!
		name, err := unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder().String(string(b))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Couldn't read your name. Are you sure it was utf-16?")

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello, %s\n", name)
	})

	fmt.Println("listening on port 8080")
	http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}
