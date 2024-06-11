package demo

import "github.com/go-chi/chi/v5"

func Routes(r *chi.Mux, handle Handle) {
	r.Mount("/api", Handler(&handle))
}
