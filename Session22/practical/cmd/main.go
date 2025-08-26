package main

import (
	"fmt"
	"net/http"
	"practical/internal/handlers"
	"practical/internal/middlewares"
)

func main() {

	http.Handle("/", middlewares.LoggMiddleware(handlers.ListFiles()))
	http.Handle("/files", middlewares.AccessMiddlewares(middlewares.LoggMiddleware(handlers.ShowFile())))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
