package go_double

type IProductRepository interface {
	FindById(id string) *Product
}
