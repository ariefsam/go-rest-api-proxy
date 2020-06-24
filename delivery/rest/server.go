package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/request", RequestHandler).Methods("GET", "POST", "PUT", "DELETE", "SET")

	handler := cors.Default().Handler(r)
	s := &http.Server{
		Addr:           ":8185",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
