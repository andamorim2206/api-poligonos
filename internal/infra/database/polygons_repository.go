package database

import (
	"database/sql"
	"fmt"
)

type PolygonsRepository struct {
	Db *sql.DB
}

func NewPolygonsRepository(db *sql.DB) *PolygonsRepository {
	return &PolygonsRepository{
		Db: db,
	}
}

func(r *PolygonsRepository) Save() error {
	_, err := r.Db.Exec("INSERT INTO poligono(id, tipo, valor_perimetro, valor_area) VALUES(003, 'triangulo', 1011, 513)");

	if err != nil {
		return err
	}

	fmt.Println("Inserção feita com Sucesso");

	return nil
}
