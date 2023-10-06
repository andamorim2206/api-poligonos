package entity

import (
	"errors"
)

const RECTANGLE string = "retangulo"
const TRIANGLE string = "triangulo"

type Perimeter struct {
	ID            string
	TypePolygons  string
	Perimeter     float64
	FirstMesuare  float64
	SecondMesuare float64
	ThirdMesuare  float64
	FourthMesuare float64
}

func NewPerimeter(params Perimeter) (*Perimeter, error) {
	perimeter := &Perimeter{
		ID:              params.ID,
		TypePolygons:    params.TypePolygons,
		FirstMesuare:	 params.FirstMesuare,	
		SecondMesuare:	 params.SecondMesuare,
		ThirdMesuare:    params.ThirdMesuare,
		FourthMesuare:   params.FourthMesuare,
	}

	err := perimeter.Validate()

	if err != nil {
		return nil, err
	}

	return perimeter, nil
}

func (p *Perimeter) Validate() error {
	if p.ID == "" {
		return errors.New("id is required")
	}

	if p.TypePolygons != RECTANGLE && p.TypePolygons != TRIANGLE {
		return errors.New("type must be retangulo or triangulo")
	}

	if p.TypePolygons == RECTANGLE && p.FirstMesuare == 0 {
		return errors.New("insert first Mesuare")
	}

	if p.TypePolygons == RECTANGLE && p.SecondMesuare == 0 {
		return errors.New("insert second Mesuare")
	}

	if p.TypePolygons == RECTANGLE && p.ThirdMesuare == 0 {
		return errors.New("insert third Mesuare")
	}

	if p.TypePolygons == RECTANGLE && p.FourthMesuare == 0 {
		return errors.New("insert fourth Mesuare")
	}

	if p.TypePolygons == TRIANGLE && p.FirstMesuare == 0 {
		return errors.New("insert first Mesuare")
	}

	if p.TypePolygons == TRIANGLE && p.SecondMesuare == 0 {
		return errors.New("insert second Mesuare")
	}
	
	if p.TypePolygons == TRIANGLE && p.ThirdMesuare == 0 {
		return errors.New("insert third Mesuare")
	}

	if p.TypePolygons == TRIANGLE && p.FourthMesuare > 0 {
		return errors.New("there is no need Fourth mesuare in Triangle")
	}

	return nil
}

func (p *Perimeter) CalculatePerimeter() error {
	p.Perimeter = p.FirstMesuare + p.SecondMesuare + p.ThirdMesuare + p.FourthMesuare
	return nil
}
