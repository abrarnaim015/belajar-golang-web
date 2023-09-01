package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface {} {
		"Title": "Template Action If",
		"Name": "Naim",
	})
}

func TestTemplateActionIf(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionIf(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{} {
		"Title": "Template Action Operator",
		"FinalValue": 50,
	})
}

func TestTemplateActionOperator(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionOperator(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{} {
		"Title": "Template Action Range",
		"Hobbies": []string {
			"Mambaca", "Menulis", "Bercerita",
		},
	})
}

func TestTemplateActionRange(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionRange(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TempalateDataWith(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(w, "address.gohtml", map[string]interface {} {
		"Title": "Template Action With",
		"Name": "Naim",
		"Address": map[string]interface{} {
			"Street": "Jl. Tubagus rangin",
			"City": "Bandung",
		},
	})
}

func TestTemplateActionWith(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TempalateDataWith(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}