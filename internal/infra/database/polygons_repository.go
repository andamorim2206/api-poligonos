package database

import (
	"database/sql"
	"fmt"

	"github.com/andamorim2206/api-poligonos/internal/entity"
)

type PolygonsRepository struct {
	Db *sql.DB
}

func NewPolygonsRepository(db *sql.DB) *PolygonsRepository {
	return &PolygonsRepository{
		Db: db,
	}
}

func(r *PolygonsRepository) Save(polygons  *entity.Polygons) error {
	_, err := r.Db.Exec("INSERT INTO poligono(id, tipo, valor_perimetro, valor_area) VALUES(?, ?, ?, ?)", 
	polygons.ID, polygons.TypePolygons, polygons.Perimeter, polygons.Area);

	if err != nil {
		return err
	}

	fmt.Println("Inserção feita com Sucesso");

	return nil
}
