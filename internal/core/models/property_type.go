package models

import "fmt"

// DisplayTranslations representa las traducciones para los nombres de campos
type DisplayTranslations struct {
	ES string `json:"es"`
	EN string `json:"en"`
	PT string `json:"pt"`
	FR string `json:"fr"`
}

// SchemaNodeValidator representa las validaciones para un campo del esquema
type SchemaNodeValidator struct {
	Required  bool     `json:"required"`
	Validator []string `json:"validator,omitempty"`
	Min       *int     `json:"min,omitempty"`
	Max       *int     `json:"max,omitempty"`
}

// SchemaNodeVisualization representa las opciones de visualización de un campo
type SchemaNodeVisualization struct {
	DisplayAs string `json:"displayAs,omitempty"`
}

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

// PropertyTypeCreate representa la estructura para crear un tipo de propiedad
type PropertyTypeCreate struct {
	Code                      string       `json:"code"`
	Display                   string       `json:"display"`
	PropertyImportPermissions []string     `json:"propertyImportPermissions,omitempty"`
	ViewPermissions           []string     `json:"viewPermissions,omitempty"`
	SchemaNodes               []SchemaNode `json:"schemaNodes"`
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
