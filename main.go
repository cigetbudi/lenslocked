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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8") //DEFAULT
	fmt.Fprint(w, "Q: Apakah ada versi gratis?<br />A: Betul ada, trial 30 hari<br /><br />Q: Mantab tidak?<br />A: Pasti mantab sekali, hubungi kami di admin@mail.com")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Halaman tidak dapat ditemukan", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Halaman tidak dapat ditemukan", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Mengaktifkan server di :3000...")
	http.ListenAndServe(":3000", router)
}
