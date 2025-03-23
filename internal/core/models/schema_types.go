package models

// BasicTypes define los tipos básicos permitidos en el esquema
const (
	// Tipos primitivos
	BasicTypeString  = "string"
	BasicTypeNumber  = "number"
	BasicTypeBoolean = "boolean"
	BasicTypeDate    = "date"

	// Tipos de referencia
	BasicTypeCOTProperty = "COTProperty"
	BasicTypeCOTUser     = "COTUser"

	// Tipos de archivo y enlace
	BasicTypeFile = "file"
	BasicTypeLink = "link"
)

// FileSubTypes define los subtipos permitidos para archivos
const (
	FileSubTypeDocument = "document"
	FileSubTypeImage    = "image"
	FileSubTypeVideo    = "video"
)

// LinkDisplayTypes define los tipos de visualización para enlaces
const (
	LinkDisplayEmbedded = "embedded"
	LinkDisplayInternal = "internal"
	LinkDisplayExternal = "external"
)

// IsValidBasicType verifica si un tipo básico es válido
func IsValidBasicType(basicType string) bool {
	switch basicType {
	case BasicTypeString, BasicTypeNumber, BasicTypeBoolean, BasicTypeDate,
		BasicTypeCOTProperty, BasicTypeCOTUser,
		BasicTypeFile, BasicTypeLink:
		return true
	default:
		return false
	}
}

// IsValidFileSubType verifica si un subtipo de archivo es válido
func IsValidFileSubType(subType string) bool {
	switch subType {
	case FileSubTypeDocument, FileSubTypeImage, FileSubTypeVideo:
		return true
	default:
		return false
	}
}

// IsValidLinkDisplay verifica si un tipo de visualización de enlace es válido
func IsValidLinkDisplay(displayType string) bool {
	switch displayType {
	case LinkDisplayEmbedded, LinkDisplayInternal, LinkDisplayExternal:
		return true
	default:
		return false
	}
}
