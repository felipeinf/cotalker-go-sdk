package cotalker

import (
	"github.com/felipeinf/cotalker_sdk-go/internal/core/cotalker"
	"github.com/felipeinf/cotalker_sdk-go/pkg/dto"
	"github.com/felipeinf/cotalker_sdk-go/pkg/models"
)

// GetProperties obtiene el servicio de propiedades del cliente
func (c *Client) getProperties() cotalker.PropertyService {
	return c.api.Properties
}

// GetAllFromCollection obtiene todos los elementos de una colección específica
func (c *Client) GetAllFromCollection(collectionCode string) ([]map[string]interface{}, error) {
	return c.getProperties().GetAllFromPropertyType(collectionCode)
}

// GetCollection obtiene todos los elementos de una colección
func (c *Client) GetAllFromCollectionPaginated(collectionCode string, params map[string]string) ([]map[string]interface{}, error) {
	return c.getProperties().GetAll(collectionCode, params)
}

// GetCollectionElement obtiene un elemento de una colección por su ID
func (c *Client) GetCollectionElement(collectionCode, id string) (map[string]interface{}, error) {
	return c.getProperties().GetByID(collectionCode, id)
}

// CreateCollectionElement crea un nuevo elemento en una colección
func (c *Client) CreateCollectionElement(collectionCode string, data map[string]interface{}) (map[string]interface{}, error) {
	return c.getProperties().Create(collectionCode, data)
}

// UpdateCollectionElement actualiza un elemento existente en una colección
func (c *Client) UpdateCollectionElement(collectionCode, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return c.getProperties().Update(collectionCode, id, data)
}

// DeleteCollectionElement elimina un elemento de una colección
func (c *Client) DeleteCollectionElement(collectionCode, id string) error {
	return c.getProperties().Delete(collectionCode, id)
}

// convertToSchemaNodes convierte un slice de maps a SchemaNodes
func convertToSchemaNodes(fields []map[string]interface{}) []models.SchemaNode {
	nodes := make([]models.SchemaNode, 0, len(fields))

	for _, field := range fields {
		node := models.SchemaNode{
			Key:       field["key"].(string),
			Display:   field["display"].(string),
			Weight:    0,
			BasicType: field["basicType"].(string),
			IsArray:   false,
			IsActive:  true,
			Validators: models.SchemaNodeValidator{
				Required: field["required"].(bool),
			},
			DisplayTranslations: models.DisplayTranslations{},
		}

		// Manejar campos opcionales
		if description, ok := field["description"].(string); ok {
			node.Description = description
		}

		if isArray, ok := field["isArray"].(bool); ok {
			node.IsArray = isArray
		}

		if subType, ok := field["subType"].(string); ok {
			node.SubType = subType
		}

		if options, ok := field["options"].([]string); ok {
			node.Validators.Validator = options
		}

		if visualization, ok := field["visualization"].(map[string]interface{}); ok {
			if displayAs, ok := visualization["displayAs"].(string); ok {
				node.Visualization = models.SchemaNodeVisualization{
					DisplayAs: displayAs,
				}
			}
		}

		nodes = append(nodes, node)
	}

	return nodes
}

// CreateCollection crea una nueva colección
func (c *Client) CreateCollection(code, display string, fields []models.SchemaNode) (map[string]interface{}, error) {
	propertyType := dto.PropertyTypeCreate{
		Code:        code,
		Display:     display,
		SchemaNodes: fields,
	}

	return c.api.PropertyTypes.Create(propertyType)
}

// GetCollections obtiene todas las colecciones disponibles
func (c *Client) GetCollections() ([]map[string]interface{}, error) {
	return c.api.PropertyTypes.GetAll()
}
