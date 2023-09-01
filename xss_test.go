package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request)  {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string] interface {} {
		"Title": "Template Auto Escape",
		"Body": "<p>Ini Adalah Body <script>alert('Anda di Heck')</script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateAutoEscape(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T)  {
	server := http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request)  {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string] interface {} {
		"Title": "Template Auto Escape",
		"Body": template.HTML("<p>Ini Adalah Body</p>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T)  {
	server := http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}