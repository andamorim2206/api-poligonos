package entity

import "errors"

const RECTANGLE string = "retangulo"
const TRIANGLE string = "triangulo"

type Polygons struct {
	ID           string
	TypePolygons string
	Perimeter    float64
	Area         float64
}

func NewPolygons(id string, typePolygons string) (*Polygons, error) {
	polygons := &Polygons{
		ID:           id,
		TypePolygons: typePolygons,
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

	if p.TypePolygons != RECTANGLE || p.TypePolygons != TRIANGLE {
		return errors.New("type must be rentangulo or triangulo")
	}

	return nil
}

func (p *Polygons) CalculatePerimeter() error {
	p.Perimeter = 1 + 2
	return nil
}
