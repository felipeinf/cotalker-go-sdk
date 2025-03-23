package cotalker

import (
	"github.com/felipeinf/cotalker_sdk-go/internal/core/config"
	"github.com/felipeinf/cotalker_sdk-go/internal/core/cotalker"
	"github.com/felipeinf/cotalker_sdk-go/internal/core/models"
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

// GetProperty obtiene todos los elementos de un tipo de propiedad
func (c *Client) GetProperty(propertyType string, params map[string]string) ([]map[string]interface{}, error) {
	return c.api.Properties.GetAll(propertyType, params)
}

// GetPropertyItem obtiene un elemento de una propiedad por su ID
func (c *Client) GetPropertyItem(propertyType, id string) (map[string]interface{}, error) {
	return c.api.Properties.GetByID(propertyType, id)
}

// GetAllFromPropertyType obtiene todos los elementos de un tipo de propiedad específico
func (c *Client) GetAllFromPropertyType(propertyType string) ([]map[string]interface{}, error) {
	return c.api.Properties.GetAllFromPropertyType(propertyType)
}

// CreatePropertyItem crea un nuevo elemento en un tipo de propiedad
func (c *Client) CreatePropertyItem(propertyType string, data map[string]interface{}) (map[string]interface{}, error) {
	return c.api.Properties.Create(propertyType, data)
}

// UpdatePropertyItem actualiza un elemento existente en un tipo de propiedad
func (c *Client) UpdatePropertyItem(propertyType, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return c.api.Properties.Update(propertyType, id, data)
}

// DeletePropertyItem elimina un elemento de un tipo de propiedad
func (c *Client) DeletePropertyItem(propertyType, id string) error {
	return c.api.Properties.Delete(propertyType, id)
}

// CreatePropertyType crea un nuevo tipo de propiedad (colección)
func (c *Client) CreatePropertyType(code, display string, schemaNodes []models.SchemaNode) (map[string]interface{}, error) {
	propertyType := models.PropertyTypeCreate{
		Code:        code,
		Display:     display,
		SchemaNodes: schemaNodes,
	}

	return c.api.PropertyTypes.Create(propertyType)
}

// GetPropertyTypes obtiene todos los tipos de propiedades disponibles
func (c *Client) GetPropertyTypes() ([]map[string]interface{}, error) {
	return c.api.PropertyTypes.GetAll()
}

// GetPropertyTypeByCode obtiene un tipo de propiedad por su código
func (c *Client) GetPropertyTypeByCode(code string) (map[string]interface{}, error) {
	return c.api.PropertyTypes.GetByCode(code)
}
