package usecase

import "github.com/andamorim2206/api-poligonos/internal/entity"

type PolygonsInput struct {
	ID           string `json:"id"`
	TypePolygons string `json:"typePolygons"`
}

type PolygonsOutput struct {
	ID           string
	TypePolygons string
	Perimeter    float64
}

type CalculatePolygons struct {
	PolygonsRepository entity.PolygonsRepositoryInterface
}

func NewCalculatePolygons(polygonsRepository entity.PolygonsRepositoryInterface) *CalculatePolygons {
	return &CalculatePolygons{
		PolygonsRepository: polygonsRepository,
	}
}

func (c *CalculatePolygons) Execute(input PolygonsInput) (*PolygonsOutput, error) {
	polygons, err := entity.NewPolygons(input.ID, input.TypePolygons)
	if err != nil {
		return nil, err
	}

	err = polygons.CalculatePerimeter()
	if err != nil {
		return nil, err
	}
	err = c.PolygonsRepository.Save(polygons)
	if err != nil {
		return nil, err
	}

	return &PolygonsOutput{
		ID:           polygons.ID,
		TypePolygons: polygons.TypePolygons,
		Perimeter:    polygons.Perimeter,
	}, nil
}
