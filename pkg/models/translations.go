package models

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
