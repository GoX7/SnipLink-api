package controlers

import "github.com/go-chi/chi"

type Handler interface {
	Register(r *chi.Mux)
}
