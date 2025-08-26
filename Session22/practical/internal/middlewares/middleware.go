package middlewares

import (
	"log"
	"net/http"
)

var access = map[string][]string{
	"vicky": {"file1.txt", "file2.txt"},
	"alice": {"file3.txt"},
}

func LoggMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s, \n", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func AccessMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qFile := r.URL.Query().Get("file")
		user := r.URL.Query().Get("user")
		if qFile == "" || user == "" {
			http.Error(w, "Please pass file and user in query", http.StatusBadRequest)
			return
		}
		files, ok := access[user]
		if !ok {
			http.Error(w, "You are not authorized to access this file.", http.StatusForbidden)
			return
		}
		for _, file := range files {
			if file == qFile {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "You are not authorized to access this file.", http.StatusForbidden)

	})
}
