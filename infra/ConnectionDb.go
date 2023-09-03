package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "dbapp.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := `INSERT INTO poligono(id, tipo, valor_perimetro, valor_area) VALUES(002, 'triangulo', 1003, 512)`;

	_, err = db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserção feita com Sucesso")
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}