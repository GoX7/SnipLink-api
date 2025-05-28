package controlers

import (
	"api/internal/logger"
	"api/internal/sqlite"
	mwlogger "api/pgk/mw_logger"
	"api/pgk/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/validator"
)

type Handlers struct {
	logs *logger.Logs
	db   *sqlite.Database
}

func NewHand(logs *logger.Logs, db *sqlite.Database) Handler {
	return &Handlers{logs: logs, db: db}
}

func (h *Handlers) Register(r *chi.Mux) {
	h.logs.Server.Debug("Register handler...")

	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(mwlogger.New(h.logs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response.NewOK())
	})
	r.Post("/l", h.SetLink)
	r.Get("/l/{alias}", h.GetLink)
}

func (h *Handlers) GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	alias := chi.URLParam(r, "alias")
	var resp response.Response

	link, err := h.db.GetLink(alias)
	if err != nil {
		switch err.Error() {
		case "G1":
			resp = response.NewError("Error database")
		case "G2":
			resp = response.NewError("Error scanner")
		case "G3":
			resp = response.NewError("Not found")
		default:
			resp = response.NewError("Unkcown error")
		}
	} else {
		resp = response.NewOkLink(link)
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handlers) SetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var link response.Request
	var resp response.Response
	json.NewDecoder(r.Body).Decode(&link)
	err := validator.New().Struct(link)
	if err != nil {
		json.NewEncoder(w).Encode(response.NewError("Validate error"))
		return
	}

	alias, err := h.db.SetLink(link.Link)
	if err != nil {
		switch err.Error() {
		case "S1":
			resp = response.NewError("Error database")
		default:
			resp = response.NewError("Unkcown error")
		}
	} else {
		resp = response.NewOkLink(alias)
	}

	json.NewEncoder(w).Encode(resp)
}
