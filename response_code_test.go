package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request)  {
	name := r.URL.Query().Get("name")

	if name == "" {
		// w.WriteHeader(400) // Bad req
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprint(w, "Name is empty")
	} else {
		// w.WriteHeader(200)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hi %s", name)
	}
}

func TestResponseCodeSuccess(t *testing.T)  {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/?name=Naim", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.Status) // 200 OK
	fmt.Println(res.StatusCode) // 200
}

func TestResponseCodeInvalid(t *testing.T)  {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.Status) // 400 Bad Request
	fmt.Println(res.StatusCode) // 400
}