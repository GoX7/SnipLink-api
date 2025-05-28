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

func NewHand(logs *logger.Logs, db *sqlite.Database) Handler { //Loading struct handlers
	return &Handlers{logs: logs, db: db}
}

func (h *Handlers) Register(r *chi.Mux) { //Register endpoint and use lib
	h.logs.Server.Debug("Register handler...")

	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(mwlogger.New(h.logs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { //Endpoint for test connection
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response.NewOK())
	})
	r.Post("/l", h.SetLink)        //Endpoint for register alias
	r.Get("/l/{alias}", h.GetLink) //Endpoint for search link
}

func (h *Handlers) GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	alias := chi.URLParam(r, "alias")

	var resp response.Response
	var code int

	link, err := h.db.GetLink(alias) //Getting link on database
	if err != nil {
		switch err.Error() {
		case "G1": //Handling error
			code = 500
			resp = response.NewError("Error database")
		case "G2":
			code = 500
			resp = response.NewError("Error scanner")
		case "G3":
			code = 404
			resp = response.NewError("Not found")
		default:
			code = 404
			resp = response.NewError("Unkcown error")
		}
	} else {
		code = 200
		resp = response.NewOkLink(link)
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp) //Encoding and send json
}

func (h *Handlers) SetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var code int
	var link response.Request
	var resp response.Response

	json.NewDecoder(r.Body).Decode(&link)
	err := validator.New().Struct(link) //Write link on database
	if err != nil {
		code = 401
		resp = response.NewError("Validate error")
	} else {
		alias, err := h.db.SetLink(link.Link)
		if err != nil {
			switch err.Error() {
			case "S1": //Handling error
				code = 500
				resp = response.NewError("Error database")
			default:
				code = 500
				resp = response.NewError("Unkcown error")
			}
		} else {
			code = 200
			resp = response.NewOkLink(alias)
		}
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp) //Encoding and send json
}
