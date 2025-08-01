package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	}))
	http.ListenAndServe(":8080", nil)
	err := errors.New("some error")
	fmt.Println(err)
}
