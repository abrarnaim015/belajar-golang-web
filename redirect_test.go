package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request)  {
	// jika ingin memasuka logic masukan di sini

	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect) // 307
}

func RedirectOut(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "https://www.google.com", http.StatusTemporaryRedirect) // 307
}

func TestRedirect(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-to-google", RedirectOut)

	server := http.Server {
		Addr: ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}