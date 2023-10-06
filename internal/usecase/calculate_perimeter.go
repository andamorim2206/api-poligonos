package usecase

import (
	"github.com/andamorim2206/api-poligonos/internal/entity"
)

type PerimeterInput struct {
	ID            string  `json:"id"`
	TypePolygons  string  `json:"tipo_poligono"`
	FirstMesuare  float64 `json:"primeira_medida"`
	SecondMesuare float64 `json:"segunda_medida"`
	ThirdMesuare  float64 `json:"terceira_medida"`
	FourthMesuare float64 `json:"quarta_medida"`
}

type PerimeterOutput struct {
	ID           string
	TypePolygons string
	Perimeter    float64
}

type CalculatePerimeter struct {
	PerimeterRepository entity.PerimeterRepositoryInterface
}

func NewCalculatePerimeter(perimeterRepository entity.PerimeterRepositoryInterface) *CalculatePerimeter {
	return &CalculatePerimeter{
		PerimeterRepository: perimeterRepository,
	}
}

func (c *CalculatePerimeter) Execute(input PerimeterInput) (*PerimeterOutput, error) {
	params := entity.Perimeter{
		ID:              input.ID,
		TypePolygons:    input.TypePolygons,
		FirstMesuare:    input.FirstMesuare,
		SecondMesuare:   input.SecondMesuare,
		ThirdMesuare:    input.ThirdMesuare,
		FourthMesuare:   input.FourthMesuare,
	}

	perimeter, err := entity.NewPerimeter(params)
	if err != nil {
		return nil, err
	}

	err = perimeter.CalculatePerimeter()
	if err != nil {
		return nil, err
	}
	err = c.PerimeterRepository.Save(perimeter)
	if err != nil {
		return nil, err
	}

	return &PerimeterOutput{
		ID:           perimeter.ID,
		TypePolygons: perimeter.TypePolygons,
		Perimeter:    perimeter.Perimeter,
	}, nil
}
