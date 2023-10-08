package entity

import (
	"errors"
)

const RECTANGLEAREA string = "retangulo"
const TRIANGLEAREA string = "triangulo"

type Area struct {
	ID            string
	TypePolygons  string
	Area          float64
	FirstMesuare  float64
	SecondMesuare float64
}

func NewArea(params Area) (*Area, error) {
	area := &Area{
		ID:            params.ID,
		TypePolygons:  params.TypePolygons,
		FirstMesuare:  params.FirstMesuare,
		SecondMesuare: params.SecondMesuare,
	}

	err := area.Validate()

	if err != nil {
		return nil, err
	}

	return area, nil
}

func (p *Area) Validate() error {
	if p.ID == "" {
		return errors.New("id is required")
	}

	if p.TypePolygons != RECTANGLEAREA && p.TypePolygons != TRIANGLEAREA {
		return errors.New("type must be retangulo or triangulo")
	}

	if p.TypePolygons == RECTANGLEAREA && p.FirstMesuare == 0 {
		return errors.New("insert first Mesuare")
	}

	if p.TypePolygons == RECTANGLEAREA && p.SecondMesuare == 0 {
		return errors.New("insert second Mesuare")
	}

	if p.TypePolygons == TRIANGLEAREA && p.FirstMesuare == 0 {
		return errors.New("insert first Mesuare")
	}

	if p.TypePolygons == TRIANGLEAREA && p.SecondMesuare == 0 {
		return errors.New("insert second Mesuare")
	}

	return nil
}

func (p *Area) CalculateArea() error {
	p.Area = p.FirstMesuare * p.SecondMesuare
	return nil
}
