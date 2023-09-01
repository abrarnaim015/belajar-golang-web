package belajargolangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request)  {
	err := myTemplates.ExecuteTemplate(w, "upload-file", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request)  {
	// r.ParseMultipartForm(32 << 20) <- default 32 MB if wanna upper can change to (100 << 20) so now is max 100 MB max file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload-success", map[string] interface {} {
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFile(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server {
		Addr: ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/big_man.png
var uploadFileTest []byte
func TestUploadFIle(t *testing.T)  {
	body := new(bytes.Buffer)

	w := multipart.NewWriter(body)
	w.WriteField("name", "Abrar Naim")
	file, _ := w.CreateFormFile("file", "UNITTEST-UPLOAD-FILE.png")
	file.Write(uploadFileTest)
	w.Close()

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rec := httptest.NewRecorder()

	Upload(rec, req)

	bodyResponse, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(bodyResponse))
}