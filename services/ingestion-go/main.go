package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Simulamos Redis en memoria para que el demo corra sin configurar DB externa
var totalRecuperado float64 = 0.0
var modoGuerra = false 

func main() {
	mux := http.NewServeMux()

	// Recibe eventos de la IA (Robo/Merma)
	mux.HandleFunc("/v1/ingest", func(w http.ResponseWriter, r *http.Request) {
		ahorro := 450.50 // Simulación de valor recuperado
		totalRecuperado += ahorro
		fmt.Printf("💰 IA DETECTÓ EVENTO: +$%.2f | Total: $%.2f\n", ahorro, totalRecuperado)
		w.Write([]byte("✅ Evento procesado"))
	})

	// Consulta de ROI para el Cliente
	mux.HandleFunc("/v1/recovery/total", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "%.2f", totalRecuperado)
	})

	// Switch de Seguridad (Fase 2)
	mux.HandleFunc("/v1/security/toggle", func(w http.ResponseWriter, r *http.Request) {
		modoGuerra = !modoGuerra
		fmt.Printf("🛡️ CAMBIO DE SEGURIDAD: Modo Guerra = %v\n", modoGuerra)
		w.Write([]byte(fmt.Sprintf("Modo Guerra: %v", modoGuerra)))
	})

	fmt.Println("🌀 CORE-COLD ACTIVADO | AUDITORÍA EN VIVO")
	server := &http.Server{
		Addr: ":8080",
		Handler: securityMiddleware(mux),
	}
	log.Fatal(server.ListenAndServe())
}

func securityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if modoGuerra {
			w.Header().Set("X-Security", "10s-ROTATION")
		} else {
			w.Header().Set("X-Security", "35m-NORMAL")
		}
		next.ServeHTTP(w, r)
	})
}
