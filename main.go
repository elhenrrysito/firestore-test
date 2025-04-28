package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time" // Para el ejemplo de uso

	"cloud.google.com/go/firestore"
	// "google.golang.org/api/option" // Descomentar si necesitas opciones explícitas (ej: credenciales)
)

// NewFirestoreClient crea un nuevo cliente de Firestore.
// Se conecta automáticamente al emulador si la variable de entorno
// FIRESTORE_EMULATOR_HOST está configurada. De lo contrario, se conecta
// al servicio real usando las credenciales predeterminadas de la aplicación (ADC)
// u otras opciones configuradas.
func NewFirestoreClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	emulatorHost := "localhost:8080"

	// Preparamos las opciones del cliente (generalmente no se necesitan para el emulador)
	// opts := []option.ClientOption{} // Inicialmente vacío

	if emulatorHost != "" {
		fmt.Printf("Variable FIRESTORE_EMULATOR_HOST detectada: %s. Conectando al emulador...\n", emulatorHost)
		// La biblioteca cliente de Go detecta FIRESTORE_EMULATOR_HOST automáticamente
		// y configura la conexión para usar el emulador sin requerir credenciales reales.
		// Por lo general, no necesitas pasar 'option.WithoutAuthentication()' explícitamente.
		// Un ID de proyecto ficticio es suficiente para el emulador si no se proporciona uno.
		if projectID == "" {
			projectID = "local-emulator-project"
			fmt.Printf("Usando ID de proyecto ficticio para el emulador: %s\n", projectID)
		}
	} else {
		fmt.Println("Variable FIRESTORE_EMULATOR_HOST no detectada. Conectando al servicio real de Firestore...")
		// Para el servicio real, el projectID es generalmente necesario si no se infiere de ADC.
		if projectID == "" {
			// Intenta obtenerlo de variables de entorno comunes si está vacío
			projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
			if projectID == "" {
				projectID = os.Getenv("GCP_PROJECT") // Otra variable común
			}
			// Si sigue vacío, NewClient podría fallar si no puede inferirlo
			if projectID == "" {
				log.Println("Advertencia: projectID está vacío. La conexión al servicio real podría fallar si no se puede inferir de ADC o del entorno.")
			} else {
				fmt.Printf("Usando projectID inferido del entorno: %s\n", projectID)
			}
		}

		// Si NO usas ADC y necesitas credenciales explícitas (ej: archivo de cuenta de servicio):
		// credentialsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") // O ruta directa
		// if credentialsPath != "" {
		//     fmt.Printf("Usando credenciales explícitas de: %s\n", credentialsPath)
		//     opts = append(opts, option.WithCredentialsFile(credentialsPath))
		// } else {
		//     fmt.Println("Usando Application Default Credentials (ADC).")
		// }
	}

	// Crear el cliente. La biblioteca maneja la lógica del emulador internamente.
	// Pasamos las opciones (opts) si las hubiéramos configurado.
	client, err := firestore.NewClient(ctx, projectID /*, opts... */) // Pasar opts si se usan
	if err != nil {
		// Proporcionar más contexto en caso de error
		log.Printf("Error al crear cliente de Firestore para proyecto '%s': %v", projectID, err)
		if emulatorHost == "" {
			log.Println("Asegúrate de que Application Default Credentials (ADC) estén configuradas (ejecuta 'gcloud auth application-default login')")
			log.Println("o que la variable de entorno GOOGLE_APPLICATION_CREDENTIALS apunte a un archivo de clave válido.")
		}
		return nil, fmt.Errorf("firestore.NewClient falló: %w", err)
	}

	if emulatorHost != "" {
		fmt.Println("Cliente de Firestore conectado exitosamente al emulador.")
	} else {
		fmt.Println("Cliente de Firestore conectado exitosamente al servicio real.")
	}
	return client, nil
}

type Venta struct {
	ID         string  `firestore:"id"`
	Producto   string  `firestore:"producto"`
	Cantidad   int     `firestore:"cantidad"`
	Precio     float64 `firestore:"precio"`
	Fecha      string  `firestore:"fecha"` // Puedes usar un tipo de tiempo más adecuado
	ClienteID  string  `firestore:"cliente_id"`
	Total      float64 `firestore:"total"`
	MetodoPago string  `firestore:"metodo_pago"`
	// Otros campos relevantes para una venta
}

// --- Ejemplo de Uso ---
func main() {
	// Es buena práctica usar contextos con timeouts o deadlines en aplicaciones reales
	ctx := context.Background() // Contexto simple para el ejemplo

	// --- Configuración ---
	// ID de tu proyecto de Google Cloud.
	// Déjalo vacío ("") si quieres que se infiera del entorno (para servicio real)
	// o si estás usando el emulador (se usará uno ficticio).
	// Si te conectas al servicio real y ADC no puede inferir el proyecto, debes establecerlo aquí.
	var projectID string = "testing-safed-d502a" // Ej: "tu-proyecto-gcp"

	// --- Crear Cliente ---
	fmt.Println("Intentando crear el cliente de Firestore...")
	client, err := NewFirestoreClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Error fatal al crear cliente: %v", err)
	}
	// ¡Importante! Asegúrate de cerrar el cliente cuando ya no lo necesites.
	// 'defer' es útil para esto.
	defer client.Close()

	fmt.Println("\nCliente creado exitosamente. Realizando operación de prueba...")

	//// --- Operación de Ejemplo: Escribir y Leer un Documento ---
	//collectionName := "go_users_test"
	//docID := "aturing"
	//
	//docRef := client.Collection(collectionName).Doc(docID)
	//
	//// Datos a escribir
	//userData := map[string]interface{}{
	//	"first":     "Alan",
	//	"middle":    "Mathison",
	//	"last":      "Turing",
	//	"born":      1912,
	//	"timestamp": firestore.ServerTimestamp, // Usar la marca de tiempo del servidor Firestore
	//}
	//
	//// Escribir (Set reemplaza el documento si existe)
	//_, err = docRef.Set(ctx, userData)
	//if err != nil {
	//	log.Fatalf("Error al escribir el documento %s/%s: %v", collectionName, docID, err)
	//}
	//fmt.Printf("Documento '%s' escrito en la colección '%s'.\n", docID, collectionName)

	// Define los datos de la nueva venta
	nuevaVenta := Venta{
		ID:         "venta123",
		Producto:   "iPhone 15",
		Cantidad:   2,
		Precio:     1099.99,
		Fecha:      "2025-04-26",
		ClienteID:  "cliente123",
		Total:      2199.98,
		MetodoPago: "Tarjeta de Crédito",
	}

	// Define la ruta al documento donde se guardará la venta
	businessID := "apple"
	paisID := "appcl" // Usamos "chile" en lugar de "APPCL" para el ID del documento país por legibilidad
	ventasCollection := "ventas"

	// Genera un nuevo documento con un ID automático dentro de la subcolección "ventas"
	_, _, err = client.Collection(businessID).
		Doc(paisID).
		Collection(ventasCollection).
		Add(context.Background(), nuevaVenta)
	if err != nil {
		log.Fatalf("error adding document: %v\n", err)
	}

	// Esperar un instante (útil a veces con el emulador o eventual consistency)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Venta creada con  en /%s/%s/%s\n", businessID, paisID, ventasCollection)

	//// Leer el documento
	//docSnapshot, err := docRef.Get(ctx)
	//if err != nil {
	//	log.Fatalf("Error al leer el documento %s/%s: %v", collectionName, docID, err)
	//}
	//
	//if docSnapshot.Exists() {
	//	// Obtener los datos como un mapa
	//	data := docSnapshot.Data()
	//	fmt.Printf("Datos leídos del documento: %+v\n", data)
	//} else {
	//	fmt.Printf("El documento %s/%s no existe.\n", collectionName, docID)
	//}

	findByClientID(client)

	fmt.Println("\nOperación de prueba completada.")
}

func findByClientID(client *firestore.Client) {
	// Define los parámetros de la búsqueda
	businessID := "apple"
	paisID := "appcl"
	ventasCollection := "ventas"
	clienteABuscar := "cliente123"

	// Realiza la consulta para encontrar la venta por ClienteID
	//iter := client.Collection(businessID).Doc(paisID).Collection(ventasCollection).Where("ClienteID", "==", clienteABuscar).Documents(context.Background())
	iter := client.Collection(businessID).Doc(paisID).Collection(ventasCollection).Documents(context.Background())

	defer iter.Stop()

	fmt.Printf("Ventas encontradas para ClienteID '%s' en /%s/%s/%s:\n", clienteABuscar, businessID, paisID, ventasCollection)

	found := false
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == context.Canceled {
				// La consulta fue cancelada
				return
			}
			// La condición para "no encontrado" es que el iterador termine sin errores después de la primera llamada a Next()
			if !found {
				fmt.Println("No se encontraron ventas para este ClienteID.")
				return
			}
			log.Fatalf("Error al iterar sobre los documentos: %v\n", err)
			return
		}
		found = true

		var ventaEncontrada Venta
		if err := doc.DataTo(&ventaEncontrada); err != nil {
			log.Fatalf("Error al convertir datos del documento: %v\n", err)
			return
		}

		fmt.Printf("    ID: %s\n", ventaEncontrada.ID)
		fmt.Printf("    Producto: %s\n", ventaEncontrada.Producto)
		fmt.Printf("    Cantidad: %d\n", ventaEncontrada.Cantidad)
		fmt.Printf("    Precio: %.2f\n", ventaEncontrada.Precio)
		fmt.Printf("    Fecha: %s\n", ventaEncontrada.Fecha)
		fmt.Printf("    Total: %.2f\n", ventaEncontrada.Total)
		fmt.Printf("    Metodo de Pago: %s\n", ventaEncontrada.MetodoPago)
		fmt.Println("---")
	}
}
