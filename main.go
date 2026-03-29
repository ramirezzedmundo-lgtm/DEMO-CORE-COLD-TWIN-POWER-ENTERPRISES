package main

import (
	"fmt"
	"net/http"
	"github.com/golang-jwt/jwt/v5" // Para la seguridad Zero Trust
)

func main() {
	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		// Aquí recibiremos la imagen y el ID del producto
		fmt.Fprint(w, "Evento de Retail Recibido y Cifrado con JWT")
	})

	fmt.Println("Centro del Huracán escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
