package usecase

import (
	"github.com/andamorim2206/api-poligonos/internal/entity"
)

type AreaInput struct {
	ID            string  `json:"id"`
	TypePolygons  string  `json:"tipo_poligono"`
	FirstMesuare  float64 `json:"primeira_medida"`
	SecondMesuare float64 `json:"segunda_medida"`
}

type AreaOutput struct {
	ID           string
	TypePolygons string
	Area         float64
}

type CalculateArea struct {
	AreaRepository entity.AreaRepositoryInterface
}

func NewCalculateArea(areaRepository entity.AreaRepositoryInterface) *CalculateArea {
	return &CalculateArea{
		AreaRepository: areaRepository,
	}
}

func (c *CalculateArea) Execute(input AreaInput) (*AreaOutput, error) {
	params := entity.Area{
		ID:            input.ID,
		TypePolygons:  input.TypePolygons,
		FirstMesuare:  input.FirstMesuare,
		SecondMesuare: input.SecondMesuare,
	}

	area, err := entity.NewArea(params)
	if err != nil {
		return nil, err
	}

	err = area.CalculateArea()
	if err != nil {
		return nil, err
	}
	err = c.AreaRepository.Save(area)
	if err != nil {
		return nil, err
	}

	return &AreaOutput{
		ID:           area.ID,
		TypePolygons: area.TypePolygons,
		Area:         area.Area,
	}, nil
}
