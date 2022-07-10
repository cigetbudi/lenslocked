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
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Halaman tidak dapat ditemukan", http.StatusNotFound)
	}
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Halaman tidak dapat ditemukan", http.StatusNotFound)
// 	}
// }

func main() {

	fmt.Println("Mengaktifkan server di :3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
