package main

import (
	"fmt"
	"net/http"

	"github.com/cjcs19/demotodevco/app"
	"github.com/rs/cors"
)

func main() {

	r := app.GetNewRoutes()

	// Start http server

	port := fmt.Sprintf(":%d", 3001)
	fmt.Println("\n\nListening on ", port)
	if err := http.ListenAndServe(port, Cors(r)); err != nil {
		return
	}
}

//Cors function to Handle request http
func Cors(r http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	return c.Handler(r)
}
