package config

// Configuration contiene la configuración para el cliente de Cotalker
type Configuration struct {
	BaseURL     string // URL base de la API de Cotalker
	AccessToken string // Token de acceso para autenticación
	CompanyID   string // ID de la empresa
}

// NewConfiguration crea una nueva instancia de Configuration
func NewConfiguration(baseURL, accessToken, companyID string) *Configuration {
	return &Configuration{
		BaseURL:     baseURL,
		AccessToken: accessToken,
		CompanyID:   companyID,
	}
}
