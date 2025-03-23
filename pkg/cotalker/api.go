package cotalker

import (
	"github.com/felipeinf/cotalker_sdk-go/internal/core/config"
	"github.com/felipeinf/cotalker_sdk-go/internal/core/cotalker"
)

// Client es el cliente público para interactuar con la API de Cotalker
type Client struct {
	api *cotalker.API
}

// NewClient crea una nueva instancia del cliente de Cotalker
func NewClient(baseURL, accessToken, companyID string) *Client {
	configuration := config.NewConfiguration(baseURL, accessToken, companyID)
	api := cotalker.NewAPI(configuration)

	return &Client{
		api: api,
	}
}
