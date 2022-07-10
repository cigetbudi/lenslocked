package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8") //DEFAULT
	fmt.Fprint(w, "<h1>SELAMAT DATANG</h1>")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8") //DEFAULT
	fmt.Fprint(w, "<h1>Contact Page</h1><p>Untuk lebih lanjut, hubungi saya ke <a href=\"mailto:cigetbudi@gmail.com\">cigetbudi@gmail.com</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func main() {
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Mengaktifkan server di :3000...")
	http.ListenAndServe(":3000", nil)
}
