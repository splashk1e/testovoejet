package handlers

import (
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/splashk1e/jet/internal/services"
)

type Handler struct {
	service *services.ServerService
	mu      *sync.Mutex
}

func NewHandler(service *services.ServerService) *Handler {
	return &Handler{
		service: service,
		mu:      &sync.Mutex{},
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("handler get request")
	switch r.URL.Path {
	case "/getstatus":
		h.GetStatus(w, r)
	default:
		h.WrongReguest(w, r)
	}
}
func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	reponse, err := h.service.GetStatus()
	if err != nil {
		logrus.Error(err.Error())
		w.Write([]byte(err.Error()))
	}
	w.Write(reponse)
}
func (h *Handler) WrongReguest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}
