package database

import (
	"database/sql"
	"fmt"

	"github.com/andamorim2206/api-poligonos/internal/entity"
)

type AreaRepository struct {
	Db *sql.DB
}

func NewAreaRepository(db *sql.DB) *AreaRepository {
	return &AreaRepository{
		Db: db,
	}
}

func(r *AreaRepository) Save(area  *entity.Area) error {
	_, err := r.Db.Exec("INSERT INTO poligono(id, tipo, valor_perimetro, valor_area) VALUES(?, ?, NULL, ?)", 
	area.ID, area.TypePolygons, area.Area);

	if err != nil {
		return err
	}

	fmt.Println("Inserção feita com Sucesso");

	return nil
}
