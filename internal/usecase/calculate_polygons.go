package usecase

import "github.com/andamorim2206/api-poligonos/internal/entity"

type PolygonsInput struct {
	ID           string `json:"id"`
	TypePolygons string `json:"typePolygons"`
}

type PolygonsOutput struct {
	Perimeter float64
	Area      float64
}

type CalculatePolygons struct {
	PolygonsRepository entity.PolygonsRepositoryInterface
}

func NewCalculatePolygons(polygonsRepository entity.PolygonsRepositoryInterface) *CalculatePolygons {
	return &CalculatePolygons{
		PolygonsRepository: polygonsRepository,
	}
}

func (c *CalculatePolygons) Execute(input PolygonsInput) (*PolygonsOutput) {
	
}