package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HomeHandler обрабатывает корневой эндпоинт
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Page loading error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

// UploadHandler обрабатывает \upload эндпоинт
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Form parsing error", http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "File receiving error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	if handler.Size == 0 {
		http.Error(w, "File empty", http.StatusInternalServerError)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File reading error", http.StatusInternalServerError)
		return
	}
	str, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resFile, err := os.Create(time.Now().UTC().Format("20060102_150405") + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resFile.Close()
	_, err = resFile.WriteString(str)
	if err != nil {
		http.Error(w, "File writing error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, str)
}
