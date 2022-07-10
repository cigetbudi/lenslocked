package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Halaman tidak dapat ditemukan", http.StatusNotFound)
	})

	fmt.Println("Mengaktifkan server di :3000...")
	http.ListenAndServe(":3000", r)
}
