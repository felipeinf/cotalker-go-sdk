package cotalker

import (
	"encoding/json"
	"fmt"
	"strconv"

	http "github.com/felipeinf/cotalker_sdk-go/internal/core/http"
)

// PropertyService define la interfaz para el servicio de propiedades
type PropertyService interface {
	GetAll(propertyType string, params map[string]string) ([]map[string]interface{}, error)
	GetByID(propertyType string, id string) (map[string]interface{}, error)
	Create(propertyType string, data map[string]interface{}) (map[string]interface{}, error)
	Update(propertyType string, id string, data map[string]interface{}) (map[string]interface{}, error)
	Delete(propertyType string, id string) error
	GetAllFromPropertyType(propertyType string) ([]map[string]interface{}, error)
}

// Implementación del servicio de propiedades
type propertyService struct {
	client *http.Client
}

// newPropertyService crea una nueva instancia del servicio de propiedades
func newPropertyService(client *http.Client) PropertyService {
	return &propertyService{
		client: client,
	}
}

// GetAll obtiene todas las propiedades del tipo especificado con los parámetros adicionales
func (s *propertyService) GetAll(propertyType string, params map[string]string) ([]map[string]interface{}, error) {
	if params == nil {
		params = make(map[string]string)
	}

	// Usar la ruta correcta y los parámetros necesarios
	path := "/api/v2/properties"
	params["propertyTypes"] = propertyType
	if params["limit"] == "" {
		params["limit"] = "200"
	}
	if params["isActive"] == "" {
		params["isActive"] = "true"
	}

	responseBytes, err := s.client.Get(path, params)
	if err != nil {
		return nil, fmt.Errorf("error al obtener propiedades del tipo %s: %w", propertyType, err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %w", err)
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		if properties, ok := data["properties"].([]interface{}); ok {
			result := make([]map[string]interface{}, 0, len(properties))
			for _, property := range properties {
				if propertyMap, ok := property.(map[string]interface{}); ok {
					result = append(result, propertyMap)
				}
			}
			return result, nil
		}
	}

	return nil, fmt.Errorf("formato de respuesta inesperado al obtener propiedades")
}

// GetAllFromPropertyType obtiene todas las propiedades del tipo especificado
func (s *propertyService) GetAllFromPropertyType(propertyType string) ([]map[string]interface{}, error) {
	// Usar la ruta correcta y los parámetros necesarios
	path := "/api/v2/properties"
	params := map[string]string{
		"propertyTypes": propertyType,
		"limit":         "200",
		"isActive":      "true",
		"count":         "true",
	}

	page := 1
	allProperties := []map[string]interface{}{}

	for {
		params["page"] = strconv.Itoa(page)
		responseBytes, err := s.client.Get(path, params)
		if err != nil {
			return nil, fmt.Errorf("error al obtener propiedades del tipo %s: %w", propertyType, err)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(responseBytes, &response); err != nil {
			return nil, fmt.Errorf("error al deserializar la respuesta: %w", err)
		}

		if data, ok := response["data"].(map[string]interface{}); ok {
			if properties, ok := data["properties"].([]interface{}); ok {
				for _, property := range properties {
					if propertyMap, ok := property.(map[string]interface{}); ok {
						allProperties = append(allProperties, propertyMap)
					}
				}
			}

			// Verificar si hay más páginas
			if count, ok := data["count"].(float64); ok {
				totalItems := int(count)
				if len(allProperties) >= totalItems {
					break
				}
			} else {
				break
			}
		} else {
			break
		}

		page++
	}

	return allProperties, nil
}

// GetByID obtiene una propiedad específica por su ID
func (s *propertyService) GetByID(propertyType string, id string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/v2/properties/%s", id)

	responseBytes, err := s.client.Get(path, nil)
	if err != nil {
		return nil, fmt.Errorf("error al obtener propiedad %s del tipo %s: %w", id, propertyType, err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %w", err)
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		return data, nil
	}

	return nil, fmt.Errorf("formato de respuesta inesperado al obtener propiedad %s", id)
}

// Create crea una nueva propiedad
func (s *propertyService) Create(propertyType string, data map[string]interface{}) (map[string]interface{}, error) {
	// Asegurarse de que el propertyType está en los datos
	data["propertyType"] = propertyType

	path := "/api/v2/properties"

	responseBytes, err := s.client.Post(path, data)
	if err != nil {
		return nil, fmt.Errorf("error al crear propiedad del tipo %s: %w", propertyType, err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %w", err)
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		return data, nil
	}

	return nil, fmt.Errorf("formato de respuesta inesperado al crear propiedad")
}

// Update actualiza una propiedad existente
func (s *propertyService) Update(propertyType string, id string, data map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/v2/properties/%s", id)

	responseBytes, err := s.client.Put(path, data)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar propiedad %s del tipo %s: %w", id, propertyType, err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %w", err)
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		return data, nil
	}

	return nil, fmt.Errorf("formato de respuesta inesperado al actualizar propiedad %s", id)
}

// Delete elimina una propiedad
func (s *propertyService) Delete(propertyType string, id string) error {
	path := fmt.Sprintf("/api/v2/properties/%s", id)

	_, err := s.client.Delete(path)
	if err != nil {
		return fmt.Errorf("error al eliminar propiedad %s del tipo %s: %w", id, propertyType, err)
	}

	return nil
}
