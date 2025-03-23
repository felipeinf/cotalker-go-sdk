package models

import (
	"fmt"
)

// SchemaNode representa un campo en el esquema de una propiedad
type SchemaNode struct {
	Weight              int                     `json:"weight"`
	Key                 string                  `json:"key"`
	Display             string                  `json:"display"`
	DisplayTranslations DisplayTranslations     `json:"displayTranslations,omitempty"`
	Description         string                  `json:"description,omitempty"`
	BasicType           string                  `json:"basicType"`
	SubType             string                  `json:"subType,omitempty"`
	IsArray             bool                    `json:"isArray"`
	IsRequired          bool                    `json:"isRequired"`
	IsActive            bool                    `json:"isActive"`
	Validators          SchemaNodeValidator     `json:"validators,omitempty"`
	Visualization       SchemaNodeVisualization `json:"visualization,omitempty"`
}

// Validate verifica que el SchemaNode tenga valores válidos
func (s *SchemaNode) Validate() error {
	// Validar que el tipo básico sea válido
	if !IsValidBasicType(s.BasicType) {
		return fmt.Errorf("tipo básico no válido: %s", s.BasicType)
	}

	// Validar subtipo para archivos
	if s.BasicType == BasicTypeFile && s.SubType != "" {
		if !IsValidFileSubType(s.SubType) {
			return fmt.Errorf("subtipo de archivo no válido: %s", s.SubType)
		}
	}

	// Validar visualización para enlaces
	if s.BasicType == BasicTypeLink && s.Visualization.DisplayAs != "" {
		if !IsValidLinkDisplay(s.Visualization.DisplayAs) {
			return fmt.Errorf("tipo de visualización de enlace no válido: %s", s.Visualization.DisplayAs)
		}
	}

	return nil
}
