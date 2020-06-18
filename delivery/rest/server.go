package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/request", RequestHandler).Methods("GET", "POST", "PUT", "DELETE", "SET")

	s := &http.Server{
		Addr:           ":8185",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
