// main.go
package main

import (
     "fmt"
     "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   		fmt.Fprintln(w, "Hello, Go World!")
	})

	http.ListenAndServe(":8080", nil)
}
