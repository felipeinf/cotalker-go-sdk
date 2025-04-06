package models

type Property struct {
	ID             string                 `json:"_id"`
	Subproperty    []string               `json:"subproperty,omitempty"`
	Name           PropertyName           `json:"name"`
	CreatedAt      string                 `json:"createdAt"`
	ModifiedAt     string                 `json:"modifiedAt"`
	PropertyType   string                 `json:"propertyType"`
	Extra          map[string]interface{} `json:"extra,omitempty"`
	SchemaInstance map[string]interface{} `json:"schemaInstance,omitempty"`
	Owner          map[string]string      `json:"owner,omitempty"`
	IsActive       *bool                  `json:"isActive,omitempty"`
}

type PropertyName struct {
	Code    string `json:"code"`
	Display string `json:"display"`
}

type PropertyPostBody struct {
	Subproperty    []string               `json:"subproperty,omitempty"`
	Name           PropertyName           `json:"name"`
	PropertyType   string                 `json:"propertyType"`
	Extra          map[string]interface{} `json:"extra,omitempty"`
	SchemaInstance map[string]interface{} `json:"schemaInstance,omitempty"`
	Owner          map[string]string      `json:"owner,omitempty"`
	IsActive       *bool                  `json:"isActive,omitempty"`
}
