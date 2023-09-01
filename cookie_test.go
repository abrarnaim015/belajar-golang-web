package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request)  {
	cookie := new(http.Cookie)

	cookie.Name = "X-PZN-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Create Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request)  {
	cookie, err := r.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestCookie(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server {
		Addr: ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=abrar-naim", nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, req)

	cookies := rec.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Abrar-naim"
	req.AddCookie(cookie)

	rec := httptest.NewRecorder()

	GetCookie(rec, req)

	// res := rec.Result()
	// body, _ := io.ReadAll(res.Body)
	
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}