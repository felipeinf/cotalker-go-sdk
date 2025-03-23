package cotalker

import (
	"encoding/json"
	"fmt"

	"github.com/felipeinf/cotalker_sdk-go/internal/core/http"
	"github.com/felipeinf/cotalker_sdk-go/pkg/dto"
)

// PropertyTypeService define la interfaz para el servicio de tipos de propiedades
type PropertyTypeService interface {
	Create(propertyType dto.PropertyTypeCreate) (map[string]interface{}, error)
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
func (s *propertyTypeService) Create(propertyType dto.PropertyTypeCreate) (map[string]interface{}, error) {
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
	params := map[string]string{
		"limit":    "100",
		"page":     "1",
		"count":    "true",
		"isActive": "true",
		"sortBy":   "display",
		"orderBy":  "asc",
	}

	responseBytes, err := s.client.Get(path, params)
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

	// Obtener el objeto data
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("formato de respuesta inesperado: data no es un objeto")
	}

	// Obtener el array propertyTypes dentro de data
	if types, ok := data["propertyTypes"].([]interface{}); ok {
		for _, pt := range types {
			if ptMap, ok := pt.(map[string]interface{}); ok {
				propertyTypes = append(propertyTypes, ptMap)
			}
		}
	}

	return propertyTypes, nil
}
