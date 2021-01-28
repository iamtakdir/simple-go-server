package httphandler

import "net/http"

// Handler for http requests
type Handler struct {
	mux *http.ServeMux
}

// Http handler
func New(s *http.ServeMux) *Handler {
	h := Handler{s}
	h.registerRoutes()

	return &h
}

// All your routes
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.HelloWorld)
}

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Your Potato Server is running :v "))
}