package main

import (
	"fmt"
	"net/http"
)

func middleware(orginalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler
		orginalHandler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}
func handle(w http.ResponseWriter, r *http.Request ) {
	// Bussiness logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("OK"))
}

func main() {
	// HandlerFunc returns a HTTP handler
	orginalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(orginalHandler))
	http.ListenAndServe(":8000", nil)
}
