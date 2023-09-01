package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T)  {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	}

	server := http.Server {
		Addr: ":8080",
    Handler: handler,
	}

	err := server.ListenAndServe()
	if err!= nil {
		panic(err)
  }
}

func TestServeMux(t *testing.T)  {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi!")
	})

	server := http.Server {
    Addr: ":8080",
    Handler: mux,
  }

	err := server.ListenAndServe()
	if err!= nil {
    panic(err)
  }
}

func TestRequest(t *testing.T)  {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	}

	server := http.Server {
    Addr: ":8080",
    Handler: handler,
  }

	err := server.ListenAndServe()
	if err!= nil {
    panic(err)
  }
}