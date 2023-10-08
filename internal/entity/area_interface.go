package entity

type AreaRepositoryInterface interface {
	Save(area *Area) error
}