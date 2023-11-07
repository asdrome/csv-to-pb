package pocketbase

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// PocketBase represents a client for interacting with the PocketBase API.
type PocketBase struct {
	endpoint string
}

// NewPocketBase creates a new instance of PocketBase client with the specified API endpoint.
func NewPocketBase(endpoint string) *PocketBase {
	return &PocketBase{
		endpoint: endpoint,
	}
}

// Collection creates a new PocketBase client for a specific collection.
func (pb *PocketBase) Collection(collection string) *PocketBase {
	return NewPocketBase(pb.endpoint + "/api/collections/" + collection + "/records")
}

// Create sends a POST request to the PocketBase API to create a new record.
func (pb *PocketBase) Create(data map[string]interface{}) error {
	// Convertir datos a formato JSON
	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Error al convertir datos a JSON:", err)
		return err
	}

	// URL del endpoint de la colección en PocketBase
	url := pb.endpoint

	// Crear una nueva solicitud POST con el cuerpo JSON
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Error al crear la solicitud POST:", err)
		return err
	}

	// Establecer el encabezado de autorización con el token
	req.Header.Set("Authorization", "Bearer TU_TOKEN") // Reemplaza "TU_TOKEN" con tu token de autorización

	// Establecer el tipo de contenido del cuerpo de la solicitud
	req.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error al hacer la solicitud POST:", err)
		return err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)

	// Mostrar el código de estado y la respuesta del servidor
	log.Println("Código de estado:", resp.Status)
	log.Println("Respuesta del servidor:", body.String())
	return nil
}

// ReadFromCSV reads data from a CSV file and returns a slice of maps.
func ReadFromCSV(csvFilePath string) ([]map[string]interface{}, error) {
	// Abrir el archivo CSV
	file, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatalln("Error al abrir el archivo CSV:", err)
		return nil, err
	}
	defer file.Close()

	// Leer el archivo CSV
	reader := csv.NewReader(file)

	// Leer los encabezados del CSV
	headers, err := reader.Read()
	if err != nil {
		log.Fatalln("Error al leer los encabezados del archivo CSV:", err)
		return nil, err
	}

	// Obtener índices de los campos específicos
	brandIndex := -1
	imageSrcIndex := -1
	nameIndex := -1
	priceIndex := -1
	discountIndex := -1

	for i, header := range headers {
		switch header {
		case "brand":
			brandIndex = i
		case "image-src":
			imageSrcIndex = i
		case "name":
			nameIndex = i
		case "price":
			priceIndex = i
		case "discount":
			discountIndex = i
		}
	}

	// Verificar que se encontraron todos los campos
	if brandIndex == -1 || imageSrcIndex == -1 || nameIndex == -1 || priceIndex == -1 || discountIndex == -1 {
		log.Fatalln("No se encontraron todos los campos requeridos en el archivo CSV.")
		return nil, errors.New("campos insuficientes en el archivo CSV")
	}

	// Inicializar un slice de mapas para almacenar los datos
	var data []map[string]interface{}

	// Procesar los datos del CSV y asignarlos a cada fila como un mapa
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln("Error al leer datos del archivo CSV:", err)
			return nil, err
		}

		// Crear un mapa para almacenar los datos de esta fila
		rowData := make(map[string]interface{})

		// Asignar los valores de los campos específicos al mapa
		rowData["brand"] = record[brandIndex]
		rowData["image_src"] = record[imageSrcIndex]
		rowData["name"] = record[nameIndex]
		// Convertir "price" a un número de punto flotante
		priceString := strings.Replace(record[priceIndex], "$", "", -1) // Eliminar el símbolo de dólar
		price, err := strconv.ParseFloat(priceString, 64)
		if err == nil {
			rowData["price"] = price
		} else {
			log.Println("Error al convertir price a número:", err)
			rowData["price"] = nil // Otra opción: asignar un valor predeterminado en caso de error
		}

		// Convertir "discount" a un número entero
		discountString := strings.Replace(record[discountIndex], "%", "", -1) // Eliminar el signo de porcentaje
		discount, err := strconv.Atoi(discountString)
		if err == nil {
			rowData["discount"] = discount
		} else {
			log.Println("Error al convertir discount a número:", err)
			rowData["discount"] = nil // Otra opción: asignar un valor predeterminado en caso de error
		}

		// Agregar los datos de esta fila al slice de mapas
		data = append(data, rowData)
	}

	return data, nil
}

// LoadFromCSV reads data from a CSV file and loads it into PocketBase.
func (pb *PocketBase) LoadFromCSV(csvFilePath string) error {
	// Leer datos del archivo CSV
	data, err := ReadFromCSV(csvFilePath)
	if err != nil {
		log.Fatalln("Error al leer datos del archivo CSV:", err)
		return err
	}

	// Iterar sobre cada fila de datos y enviarla a PocketBase usando el método Create
	for _, record := range data {
		err := pb.Create(record)
		if err != nil {
			log.Println("Error al enviar datos a PocketBase:", err)
			// Puedes decidir cómo manejar el error, por ejemplo, continuar con las siguientes filas o devolver el error.
		} else {
			log.Println("Datos enviados a PocketBase:", record)
		}

		// Agregar un delay de 1 segundo (1000 milisegundos) entre las peticiones
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}
