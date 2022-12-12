package go_double

import "errors"

var (
	ErrMissingArgs = errors.New("MissingArgs")
)

type ProductService struct{}

func (p *ProductService) FindById(repository IProductRepository, id string) (*Product, error) {
	if id == "" {
		return nil, ErrMissingArgs
	}

	result := repository.FindById(id)

	return result, nil

}
