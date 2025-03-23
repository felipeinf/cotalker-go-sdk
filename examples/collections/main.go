package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/felipeinf/cotalker_sdk-go/pkg/cotalker"
	"github.com/felipeinf/cotalker_sdk-go/pkg/models"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde el archivo .env local
	if err := godotenv.Load(); err != nil {
		log.Printf("Error: No se pudo cargar el archivo .env: %v", err)
	}

	// Usar las variables de entorno existentes
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		log.Fatal("Error: Se requiere una URL base (BASE_URL)")
	}

	accessToken := os.Getenv("COTALKER_TOKEN")
	if accessToken == "" {
		log.Fatal("Error: Se requiere un token de acceso (COTALKER_TOKEN)")
	}

	companyID := os.Getenv("COMPANY_ID")

	// Crear el cliente de Cotalker
	client := cotalker.NewClient(baseURL, accessToken, companyID)

	// Inicializar el generador de números aleatorios
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)

	// Definir el tipo de propiedad a crear con un número aleatorio
	collection := fmt.Sprintf("employees_%d", randomNum)

	// Definir un esquema básico para empleados
	schemaNodes := []models.SchemaNode{
		{
			Weight:    1,
			Key:       "complete_name",
			Display:   "complete_name",
			BasicType: "string",
			IsArray:   false, // default is false
			IsActive:  true,  // default is true
		},
		{
			Weight:    2,
			Key:       "salary",
			Display:   "salary",
			BasicType: "number",
		},
		{
			Weight:    3,
			Key:       "responsibilities",
			Display:   "responsibilities",
			BasicType: "string",
			IsArray:   true,
		},
	}

	// Crear la colección
	_, err := client.CreateCollection(collection, collection, schemaNodes)
	if err != nil {
		log.Fatalf("Error al crear la colección [%s]: %v", collection, err)
	}
	fmt.Printf("Colección [%s] creada exitosamente\n", collection)

	// Definir los datos del empleado
	element := map[string]interface{}{
		"name": map[string]interface{}{
			"code":    fmt.Sprintf("emp_%d", randomNum),
			"display": fmt.Sprintf("emp_%d", randomNum),
		},
		"schemaInstance": map[string]interface{}{
			"complete_name": "Juan Pérez",
			"salary":        8000,
			"responsibilities": []string{
				"Desarrollar software",
				"Mantener sistemas operativos",
				"Realizar pruebas de software",
			},
		},
	}

	// Crear el empleado
	_, err = client.CreateCollectionElement(collection, element)
	if err != nil {
		log.Fatalf("Error al crear empleado: %v", err)
	}
	fmt.Printf("Empleado creado exitosamente\n")

	// Esperar un momento para que la colección esté disponible
	time.Sleep(2 * time.Second)

	items, err := client.GetAllFromCollection(collection)
	if err != nil {
		log.Fatalf("Error al obtener empleados: %v", err)
	}

	fmt.Printf("Se encontraron %d empleados\n\n", len(items))

	// Mostrar información de cada empleado
	for _, item := range items {
		itemJSON, _ := json.MarshalIndent(item, "", "  ")
		fmt.Printf("%s\n", string(itemJSON))
	}
}
