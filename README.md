# SDK-GO para Cotalker

Este SDK proporciona una interfaz en Go para interactuar con la API de Cotalker, facilitando la gestión de colecciones y sus elementos.

## Instalación

Para usar el SDK en tu proyecto:

```bash
go get github.com/felipeinf/cotalker-go
```

## Configuración

El SDK requiere las siguientes variables de entorno:

```env
BASE_URL=https://staging.cotalker.com  # URL base de la API
COTALKER_TOKEN=tu_token_aqui          # Token de acceso
COMPANY_ID=tu_company_id              # ID de la compañía
```

También puedes configurar estas variables programáticamente al crear el cliente:

```go
client := cotalker.NewClient(baseURL, accessToken, companyID)
```

## Uso Básico

### Crear un Cliente

```go
import "github.com/felipeinf/cotalker-go/pkg/cotalker"

client := cotalker.NewClient(baseURL, accessToken, companyID)
```

### Gestión de Colecciones

```go
// Crear una colección
schemaNodes := []models.SchemaNode{
    {
        Weight:    1,
        Key:       "name",
        Display:   "Nombre",
        BasicType: models.BasicTypeString,
        IsArray:   false,
        IsActive:  true,
        Validators: models.SchemaNodeValidator{
            Required: true,
        },
    },
    // ... más nodos del esquema
}

collection, err := client.CreateCollection("employees", "Empleados", schemaNodes)

// Obtener todas las colecciones
collections, err := client.GetCollections()

// Obtener una colección específica
collection, err := client.GetCollectionByCode("employees")
```

### Gestión de Elementos de Colección

```go
// Crear un elemento
elementData := map[string]interface{}{
    "name": map[string]interface{}{
        "code":    "emp_001",
        "display": "emp_001",
    },
    "schemaInstance": map[string]interface{}{
        "name":      "Juan",
        "last_name": "Pérez",
        "salary":    8000,
        "position":  "Desarrollador",
    },
}

element, err := client.CreateCollectionElement("employees", elementData)

// Obtener todos los elementos
elements, err := client.GetAllFromCollection("employees")

// Obtener elementos con paginación
params := map[string]string{
    "page": "1",
    "limit": "100",
    "isActive": "true",
}
elements, err := client.GetAllFromCollectionPaginated("employees", params)

// Obtener un elemento específico
element, err := client.GetCollectionElement("employees", "element_id")

// Actualizar un elemento
updateData := map[string]interface{}{
    "schemaInstance": map[string]interface{}{
        "salary": 8500,
    },
}
element, err := client.UpdateCollectionElement("employees", "element_id", updateData)

// Eliminar un elemento
err := client.DeleteCollectionElement("employees", "element_id")
```

## Tipos de Datos Soportados

El SDK soporta los siguientes tipos básicos para los campos del esquema:

- `string`: Texto
- `number`: Números
- `boolean`: Valores booleanos
- `date`: Fechas
- `COTProperty`: Referencias a otras colecciones
- `COTUser`: Referencias a usuarios
- `file`: Archivos (document, image, video)
- `link`: Enlaces (embedded, internal, external)

## Ejemplos

Puedes encontrar ejemplos completos de uso en el directorio `examples/`:

- `collections/`: Ejemplo de creación y gestión de colecciones

## Desarrollo

### Estructura del Proyecto

```
go/
├── internal/      (Código privado del SDK)
│   ├── core/      (Lógica principal y modelos)
│   └── http/      (Cliente HTTP y manejo de requests)
├── pkg/           (API pública)
│   └── cotalker/  (Cliente principal y tipos exportados)
└── examples/      (Ejemplos de uso)
```

### Requisitos

- Go 1.22 o superior
- Variables de entorno configuradas (ver sección de Configuración)

### Instalación de Dependencias

```bash
go mod tidy
```
