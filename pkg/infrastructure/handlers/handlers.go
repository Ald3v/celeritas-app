package handlers

import (
	"net/http"

	"github.com/ald3v/celeritas"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	
}