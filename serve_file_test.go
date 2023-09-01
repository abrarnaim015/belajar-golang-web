package belajargolangweb

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func ServeFile(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T)  {
	server := http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceOK)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T)  {
	server := http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeFileServerEmbedTest(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "hhtp://localhost:8080?name=naim", nil)
	rec := httptest.NewRecorder()

	ServeFileEmbed(rec,req)

	body, _ := io.ReadAll(rec.Body)
	fmt.Println(string(body))
}