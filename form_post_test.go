package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// jika ingin mengambil datanya body tampa r.ParseForm() maka bisa menggunakan
	// r.PostFormValue("first_name")

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T)  {
	reqBody := strings.NewReader("first_name=Abrar&last_name=Naim")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/crateuser", reqBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	FormPost(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}