package controlers

import "github.com/go-chi/chi"

type Handler interface { //Interface for handlers
	Register(r *chi.Mux)
}
