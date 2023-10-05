package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andamorim2206/api-poligonos/internal/usecase"
	"github.com/andamorim2206/api-poligonos/internal/infra/database"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	db, err := sql.Open("sqlite3", "dbapp.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	r.HandleFunc("/calcular", CalculatePerimeter(db)).Methods("POST")
	
	fmt.Println("Servidor iniciado na porta 8000")

	http.ListenAndServe(":8000", r)
}

func CalculatePerimeter(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		polygonsRepo := database.NewPolygonsRepository(db)
		calculatePolygons := usecase.NewCalculatePolygons(polygonsRepo)

		var requestBody usecase.PolygonsInput
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := calculatePolygons.Execute(requestBody)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Perímetro calculado: %f Area Calculada: %f", output.Perimeter, output.Area)
	}
}
