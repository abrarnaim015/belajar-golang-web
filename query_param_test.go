package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, req *http.Request)  {
	name := req.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=naim", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(w http.ResponseWriter, req *http.Request)  {
	firstName := req.URL.Query().Get("first_name")
	lastName := req.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=abrar&last_name=naim", nil)
	rec := httptest.NewRecorder()

	MultipleQueryParameter(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request)  {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=abrar&name=naim", nil)
  rec := httptest.NewRecorder()

  MultipleParameterValues(rec, req)

  res := rec.Result()
  body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}