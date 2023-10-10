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

func (r *AreaRepository) FindAllArea() ([]entity.Area, error) {
	rows, err := r.Db.Query("SELECT id, tipo, valor_area FROM poligonos");
    if err != nil {
        return nil, err
    }

	defer rows.Close();

	areas := make([]entity.Area, 0)

	for rows.Next() {
		var area entity.Area
		err := rows.Scan(&area.ID, &area.TypePolygons, &area.Area)
		if err != nil {
            return nil, err
        }

		areas = append(areas, area)
	}

	return areas, nil
}