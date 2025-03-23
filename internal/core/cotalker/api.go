package cotalker

import (
	"github.com/felipeinf/cotalker_sdk-go/internal/core/config"
	httpClient "github.com/felipeinf/cotalker_sdk-go/internal/core/http"
)

// API es la interfaz principal para interactuar con la API de Cotalker
type API struct {
	client        *httpClient.Client
	configuration *config.Configuration

	// Servicio para propiedades
	Properties PropertyService

	// Servicio para tipos de propiedades
	PropertyTypes PropertyTypeService
}

// NewAPI crea una nueva instancia de la API de Cotalker
func NewAPI(configuration *config.Configuration) *API {
	client := httpClient.NewClient(configuration)

	api := &API{
		client:        client,
		configuration: configuration,
	}

	// Inicializar servicios
	api.Properties = newPropertyService(client)
	api.PropertyTypes = newPropertyTypeService(client)

	return api
}
