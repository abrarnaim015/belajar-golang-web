package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TempalateDataMap(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{} {
		"Title": "Template Data Map",
		"Name": "Naim",
		"Address": map[string]interface{} {
			"Street": "Jl. Tubagus Rangin",
		},
	})
}

func TestTemplateDataMap(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "hhtp://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TempalateDataMap(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title string
	Name string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page {
		Title: "Template Data Struct",
		Name: "Naim",
		Address: Address{
			Street: "Jl. Tubagus Rangin",
		},
	})
}

func TestTempalateDataStruct(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}