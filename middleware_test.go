package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handle http.Handler
}

type ErrorHandler struct {
	Handle http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Before Execute Handler")
	middleware.Handle.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println("RECOVER : ", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "ERROR : %s", err)
		}
	} ()

	errorHandler.Handle.ServeHTTP(w, r)
}


func TestMiddleware(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Handler Executed")
		fmt.Fprintf(w, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Foo Executed")
		fmt.Fprintf(w, "Hello Foo")
	})
	mux.HandleFunc("/panic", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Panic Executed")
		panic("Ups")
	})

	// logMiddleware := &LogMiddleware {
	// 	Handle: mux,
	// }
	// same think but below more simple <-

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handle = mux

	errorHandler := &ErrorHandler {
		Handle: logMiddleware,
	}

	server := http.Server {
		Addr: ":8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}