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

func SimpleHTML(w http.ResponseWriter, r *http.Request)  {
	templateText := `<html><body>{{.}}</body></html>`

	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }
	
	// cara lebih sederhana dari yang di atas
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	
	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)
	
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFile(rec, req)
	
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLDirectory(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateDirectory(rec, req)
	
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var tempaltes embed.FS
func TemplateEmbed(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFS(tempaltes, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLEmbed(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)
	
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}