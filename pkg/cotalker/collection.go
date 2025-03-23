package cotalker

import (
	"github.com/felipeinf/cotalker_sdk-go/internal/core/models"
)

// GetAllFromCollection obtiene todos los elementos de una colección específica
func (c *Client) GetAllFromCollection(collectionCode string) ([]map[string]interface{}, error) {
	return c.api.Properties.GetAllFromPropertyType(collectionCode)
}

// GetCollection obtiene todos los elementos de una colección
func (c *Client) GetAllFromCollectionPaginated(collectionCode string, params map[string]string) ([]map[string]interface{}, error) {
	return c.api.Properties.GetAll(collectionCode, params)
}

// GetCollectionElement obtiene un elemento de una colección por su ID
func (c *Client) GetCollectionElement(collectionCode, id string) (map[string]interface{}, error) {
	return c.api.Properties.GetByID(collectionCode, id)
}

// CreateCollectionElement crea un nuevo elemento en una colección
func (c *Client) CreateCollectionElement(collectionCode string, data map[string]interface{}) (map[string]interface{}, error) {
	return c.api.Properties.Create(collectionCode, data)
}

// UpdateCollectionElement actualiza un elemento existente en una colección
func (c *Client) UpdateCollectionElement(collectionCode, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return c.api.Properties.Update(collectionCode, id, data)
}

// DeleteCollectionElement elimina un elemento de una colección
func (c *Client) DeleteCollectionElement(collectionCode, id string) error {
	return c.api.Properties.Delete(collectionCode, id)
}

// CreateCollection crea una nueva colección
func (c *Client) CreateCollection(code, display string, schemaNodes []models.SchemaNode) (map[string]interface{}, error) {
	propertyType := models.PropertyTypeCreate{
		Code:        code,
		Display:     display,
		SchemaNodes: schemaNodes,
	}

	return c.api.PropertyTypes.Create(propertyType)
}

// GetCollections obtiene todas las colecciones disponibles
func (c *Client) GetCollections() ([]map[string]interface{}, error) {
	return c.api.PropertyTypes.GetAll()
}
