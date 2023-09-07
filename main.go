package main

import (
	"database/sql"
	"fmt"

	"github.com/andamorim2206/api-poligonos/internal/infra/database"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	db, err := sql.Open("sqlite3", "dbapp.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	polygonsRepository := database.NewPolygonsRepository(db)

	polygonsRepository.Save()

	fmt.Println("Deu Bom");
}

