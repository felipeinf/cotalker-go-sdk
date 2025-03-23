package models

type PropertyType struct {
	ID       string `json:"_id"`
	Company  string `json:"company"`
	Code     string `json:"code"`
	Display  string `json:"display"`
	IsActive bool   `json:"isActive"`
	SExtras  bool   `json:"sextras"`
	Create   struct {
		RequiredPermission []string `json:"requiredPermission"`
	} `json:"create"`
	ViewPermissions []string     `json:"viewPermissions"`
	SchemaNodes     []SchemaNode `json:"schemaNodes"`
	Visualization   struct {
		SchemaNodesGroups []interface{} `json:"schemaNodesGroups"`
	} `json:"visualization"`
	PropertyImportPermissions []string `json:"propertyImportPermissions"`
	CreatedAt                 string   `json:"createdAt"`
	ModifiedAt                string   `json:"modifiedAt"`
}
