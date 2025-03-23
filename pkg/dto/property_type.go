package dto

import (
	"fmt"

	"github.com/felipeinf/cotalker_sdk-go/pkg/models"
)

// PropertyTypeCreate representa la estructura para crear un tipo de propiedad
type PropertyTypeCreate struct {
	Code                      string              `json:"code"`
	Display                   string              `json:"display"`
	PropertyImportPermissions []string            `json:"propertyImportPermissions,omitempty"`
	ViewPermissions           []string            `json:"viewPermissions,omitempty"`
	SchemaNodes               []models.SchemaNode `json:"schemaNodes"`
}

// Validate verifica que el PropertyTypeCreate tenga valores válidos
func (p *PropertyTypeCreate) Validate() error {
	if p.Code == "" {
		return fmt.Errorf("el código es requerido")
	}
	if p.Display == "" {
		return fmt.Errorf("el display es requerido")
	}
	if len(p.SchemaNodes) == 0 {
		return fmt.Errorf("se requiere al menos un nodo en el esquema")
	}

	// Validar cada nodo del esquema
	for i, node := range p.SchemaNodes {
		if err := node.Validate(); err != nil {
			return fmt.Errorf("error en el nodo %d: %w", i+1, err)
		}
	}

	return nil
}
