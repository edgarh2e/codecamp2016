package main

import (
	"github.com/edgarh2e/codecamp2016/web/compare"
	"github.com/pressly/chi"
	"log"
	"net/http"
)

func listenAndServe(addr string) error {
	r := chi.NewRouter()

	r.Route("/compare", compare.Routes)

	return http.ListenAndServe(addr, r)
}

func main() {
	log.Fatalf("listenAndServe: %v", listenAndServe(":3000"))
}
