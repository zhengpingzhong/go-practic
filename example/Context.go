package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello2(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("Server: hello hander ended")
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalServerError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalServerError)
	}
}
func main() {
	http.HandleFunc("/hello", hello2)
	http.ListenAndServe(":8090", nil)
}
