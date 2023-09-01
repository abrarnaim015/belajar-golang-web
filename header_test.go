package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request)  {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T)  {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	RequestHeader(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func responseHeader(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("X-Power-By", "Naim")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T)  {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	responseHeader(rec, req)
	res := rec.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.Header.Get("x-power-by"))
}