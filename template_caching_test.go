package belajargolangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var cachingTemplates embed.FS
var myTemplates = template.Must(template.ParseFS(cachingTemplates, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request)  {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello Template Caching")
}

func TestTemplatecaching(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateCaching(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}