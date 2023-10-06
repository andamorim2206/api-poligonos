package entity

import (
	"errors"
	"fmt"
)

// const RECTANGLE string = "retangulo"
// const TRIANGLE string = "triangulo"

type Polygons struct {
	ID              string
	TypePolygons    string
	Perimeter       float64
	Area            float64
	TypeCalculation string
}

func NewPolygons(id string, typePolygons string, typeCalculation string) (*Polygons, error) {
	polygons := &Polygons{
		ID:           id,
		TypePolygons: typePolygons,
		TypeCalculation: typeCalculation,
	}

	err := polygons.Validate()

	if err != nil {
		return nil, err
	}

	return polygons, nil
}

func (p *Polygons) Validate() error {
	if p.ID == "" {
		return errors.New("id is required")
	}

	if p.TypePolygons != RECTANGLE && p.TypePolygons != TRIANGLE {
		return errors.New("type must be retangulo or triangulo")
	}

	return nil
}

func (p *Polygons) CalculatePerimeter() error {
	p.Perimeter = 1 + 2
	fmt.Println(p.Perimeter)
	return nil
}

func (p *Polygons) CalculateArea() error {
	p.Area = 2 * 2
	fmt.Println(p.Area)
	return nil
}
