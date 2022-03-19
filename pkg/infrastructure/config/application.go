package config

import (
	"myapp/pkg/infrastructure/handlers"
	"myapp/pkg/infrastructure/persistence"

	"github.com/ald3v/celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models   persistence.Models
}