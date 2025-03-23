package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/felipeinf/cotalker_sdk-go/internal/core/config"
)

// Client es un cliente HTTP para comunicarse con la API de Cotalker
type Client struct {
	config     *config.Configuration
	httpClient *http.Client
}

// NewClient crea una nueva instancia de Client
func NewClient(config *config.Configuration) *Client {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: time.Second * 30, // timeout por defecto de 30 segundos
		},
	}
}

// buildURL construye una URL correcta combinando la URL base y el path
func (c *Client) buildURL(path string) string {
	// Asegurarse de que baseURL no termine con barra y path comience con barra
	baseURL := strings.TrimRight(c.config.BaseURL, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return baseURL + path
}

// Get realiza una petición GET a la API de Cotalker
func (c *Client) Get(path string, params map[string]string) ([]byte, error) {
	url := c.buildURL(path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición HTTP: %w", err)
	}

	// Agregar parámetros de consulta
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// Agregar cabeceras de autenticación y de empresa
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("admin", "true")
	if c.config.CompanyID != "" {
		req.Header.Add("Company", c.config.CompanyID)
	}

	return c.do(req)
}

// Post realiza una petición POST a la API de Cotalker
func (c *Client) Post(path string, body interface{}) ([]byte, error) {
	url := c.buildURL(path)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error al serializar el cuerpo de la petición: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición HTTP: %w", err)
	}

	// Agregar cabeceras de autenticación y de empresa
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("admin", "true")
	if c.config.CompanyID != "" {
		req.Header.Add("Company", c.config.CompanyID)
	}

	return c.do(req)
}

// Put realiza una petición PUT a la API de Cotalker
func (c *Client) Put(path string, body interface{}) ([]byte, error) {
	url := c.buildURL(path)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error al serializar el cuerpo de la petición: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición HTTP: %w", err)
	}

	// Agregar cabeceras de autenticación y de empresa
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("admin", "true")
	if c.config.CompanyID != "" {
		req.Header.Add("Company", c.config.CompanyID)
	}

	return c.do(req)
}

// Delete realiza una petición DELETE a la API de Cotalker
func (c *Client) Delete(path string) ([]byte, error) {
	url := c.buildURL(path)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición HTTP: %w", err)
	}

	// Agregar cabeceras de autenticación y de empresa
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("admin", "true")
	if c.config.CompanyID != "" {
		req.Header.Add("Company", c.config.CompanyID)
	}

	return c.do(req)
}

// do ejecuta una petición HTTP y procesa la respuesta
func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la petición HTTP: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("error en la respuesta HTTP (status: %d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}
