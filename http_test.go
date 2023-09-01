package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello World!")
}

func TestHelloHandler(t *testing.T)  {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	bodyStr := string(body)
	fmt.Println(bodyStr)
}
