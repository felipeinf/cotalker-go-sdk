package models

type PropertyType struct {
	ID              string                   `json:"_id"`
	Code            string                   `json:"code"`
	Company         string                   `json:"company"`
	CreatedAt       string                   `json:"createdAt"`
	IsActive        bool                     `json:"isActive"`
	ModifiedAt      string                   `json:"modifiedAt"`
	Display         string                   `json:"display"`
	SchemaNodes     []PropertyTypeSchemaNode `json:"schemaNodes"`
	ViewPermissions []string                 `json:"viewPermissions"`
}

type PropertyTypeSchemaNode struct {
	Validators  PropertyTypeValidator `json:"validators"`
	IsArray     bool                  `json:"isArray"`
	Weight      int                   `json:"weight"`
	IsActive    bool                  `json:"isActive"`
	Key         string                `json:"key"`
	Display     string                `json:"display"`
	Description string                `json:"description"`
	BasicType   string                `json:"basicType"` // 'string'|'number'|'boolean'|'COTProperty'|'COTUser'|'date'
	SubType     string                `json:"subType"`
}

type PropertyTypeValidator struct {
	Required bool `json:"required"`
	Min      *int `json:"min,omitempty"`
	Max      *int `json:"max,omitempty"`
}
