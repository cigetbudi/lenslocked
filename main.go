package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8") //DEFAULT
	fmt.Fprint(w, "<h1>SELAMAT DATANG</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Mengaktifkan server di :3000...")
	http.ListenAndServe(":3000", nil)
}
