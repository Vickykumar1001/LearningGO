package handlers

import (
	"encoding/json"
	"net/http"
	"practical/internal/services"
)

func ListFiles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			files, err := services.ReadDirectory("./test_files")
			if err != nil {
				http.Error(w, "Server Error", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			response := map[string][]string{"files": files}
			json.NewEncoder(w).Encode(response)
			return
		}
		http.Error(w, "Method Error", http.StatusMethodNotAllowed)
	})
}

func ShowFile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			file := r.URL.Query().Get("file")
			content, err := services.ReadFileLines("./test_files/" + file)
			if err != nil {
				http.Error(w, "Server Error", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(content)
			return
		}
		http.Error(w, "Method Error", http.StatusMethodNotAllowed)
	})
}
