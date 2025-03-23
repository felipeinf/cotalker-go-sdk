package cotalker

import (
	"encoding/json"
	"fmt"

	http "github.com/felipeinf/cotalker_sdk-go/internal/core/http"
	"github.com/felipeinf/cotalker_sdk-go/internal/core/models"
)

// PropertyTypeService define la interfaz para el servicio de tipos de propiedades
type PropertyTypeService interface {
	Create(propertyType models.PropertyTypeCreate) (map[string]interface{}, error)
	GetAll() ([]map[string]interface{}, error)
}

type propertyTypeService struct {
	client *http.Client
}

// newPropertyTypeService crea una nueva instancia del servicio de tipos de propiedades
func newPropertyTypeService(client *http.Client) PropertyTypeService {
	return &propertyTypeService{
		client: client,
	}
}

// Create crea un nuevo tipo de propiedad
func (s *propertyTypeService) Create(propertyType models.PropertyTypeCreate) (map[string]interface{}, error) {
	// Validar la estructura del tipo de propiedad
	if err := propertyType.Validate(); err != nil {
		return nil, fmt.Errorf("error de validación: %w", err)
	}

	path := "/api/v2/propertytypes"

	responseBytes, err := s.client.Post(path, propertyType)
	if err != nil {
		return nil, fmt.Errorf("error al crear tipo de propiedad %s: %w", propertyType.Code, err)
	}

	// Procesar la respuesta
	var response map[string]interface{}
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %w", err)
	}

	// Extraer el objeto data de la respuesta
	if data, ok := response["data"].(map[string]interface{}); ok {
		return data, nil
	}

	return nil, fmt.Errorf("formato de respuesta inesperado al crear tipo de propiedad")
}

// GetAll obtiene todos los tipos de propiedades
func (s *propertyTypeService) GetAll() ([]map[string]interface{}, error) {
	path := "/api/v2/propertytypes"

	responseBytes, err := s.client.Get(path, nil)
	if err != nil {
		return nil, fmt.Errorf("error al obtener tipos de propiedades: %w", err)
	}

	// Procesar la respuesta
	var response map[string]interface{}
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %w", err)
	}

	// Extraer el array de datos
	var propertyTypes []map[string]interface{}

	if data, ok := response["data"].([]interface{}); ok {
		// Si data es un array
		for _, pt := range data {
			if ptMap, ok := pt.(map[string]interface{}); ok {
				propertyTypes = append(propertyTypes, ptMap)
			}
		}
	} else if data, ok := response["data"].(map[string]interface{}); ok {
		// Si data contiene un campo "properties"
		if properties, ok := data["properties"].([]interface{}); ok {
			for _, prop := range properties {
				if propMap, ok := prop.(map[string]interface{}); ok {
					propertyTypes = append(propertyTypes, propMap)
				}
			}
		}
	}

	return propertyTypes, nil
}
