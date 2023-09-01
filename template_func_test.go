package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string  {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Bubdi" }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage {
		Name: "Naim",
	})
}

func TestTemplateFunction(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateFunction(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateGlobal(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.New("GLOBAL").Parse(`{{ len .Name }}`))
	t.ExecuteTemplate(w, "GLOBAL", MyPage {
		Name: "Naim",
	})
}

func TestTemplateGlobal(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateGlobal(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionMap(w http.ResponseWriter, r *http.Request)  {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface {} {
		"upper": func (value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage {
		Name: "naim",
	})
}

func TestTemplateMap(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateFunctionMap(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request)  {
	t := template.New("FUNCTION")
	t.Funcs(map[string] interface {} {
		"sayHello" : func (value string) string {
			return "Hello " + value
		},
		"upper" : func (value string) string {
			return strings.ToUpper(value)
		},
		"add" : func (value string) string {
			return value + ". Im a Backend Golang Developer"
		},
	})

	template.Must(t.Parse(`{{ sayHello .Name | add | upper }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage {
		Name: "naim",
	})
}

func TestTemplateFunctionPipelines(t *testing.T)  {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateFunctionPipelines(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}