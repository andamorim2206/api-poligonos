package entity

type PerimeterRepositoryInterface interface {
	Save(perimeter *Perimeter) error
}