package database

import (
	"database/sql"
	"fmt"

	"github.com/andamorim2206/api-poligonos/internal/entity"
)

type PerimeterRepository struct {
	Db *sql.DB
}

func NewPerimeterRepository(db *sql.DB) *PerimeterRepository {
	return &PerimeterRepository{
		Db: db,
	}
}

func(r *PerimeterRepository) Save(perimeter  *entity.Perimeter) error {
	_, err := r.Db.Exec("INSERT INTO poligono(id, tipo, valor_perimetro, valor_area) VALUES(?, ?, ?, NULL)", 
	perimeter.ID, perimeter.TypePolygons, perimeter.Perimeter);

	if err != nil {
		return err
	}

	fmt.Println("Inserção feita com Sucesso");

	return nil
}
