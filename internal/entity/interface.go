package entity

type PolygonsRepositoryInterface interface {
	Save(polygons *Polygons) error
}