package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5" // Seguridad de grado industrial
)

// Estructura del evento de Retail (Lo que manda la cámara/sensor)
type RetailEvent struct {
	DeviceID  string  `json:"device_id"`
	ProductID string  `json:"product_id"`
	Weight    float64 `json:"weight"`
	Timestamp int64   `json:"timestamp"`
}

var jwtKey = []byte("core_cold_twin_power_secret_2026")

func main() {
	// Endpoint principal de Ingestión
	http.HandleFunc("/v1/ingest", ingestHandler)

	fmt.Println("🚀 Cold-Gateway GO - Centro del Huracán Operativo")
	fmt.Println("📡 Escuchando señales en puerto 8080...")
	
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func ingestHandler(w http.ResponseWriter, r http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 1. Simulación de Validación Zero-Trust (JWT)
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		fmt.Println("⚠️ Intento de acceso no autorizado detectado.")
		http.Error(w, "No autorizado - Zero Trust Active", http.StatusUnauthorized)
		return
	}

	// 2. Decodificar el evento de Retail
	var event RetailEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Datos corruptos", http.StatusBadRequest)
		return
	}

	// 3. Envío al "Cerebro" (Simulado para el MVP)
	fmt.Printf("📦 Evento Recibido: Producto %s desde Dispositivo %s\n", event.ProductID, event.DeviceID)
	
	// Respuesta rápida para mantener la baja latencia
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status": "procesando", "audit_id": "HASH-RT-99"}`))
}
