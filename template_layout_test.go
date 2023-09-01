package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml",
	))
	t.ExecuteTemplate(w, "layout", map[string]interface{} {
		"Title": "Template Layout",
		"Name": "Naim",
	})
}

func TestTemplateLayout(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateLayout(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}