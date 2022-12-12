package go_double

import "errors"

var (
	ErrMissingArgs = errors.New("MissingArgs")
)

type ProductService struct{}

func (p *ProductService) FindById(repository ProductRepository , id string)( *Product , error ){
	if id == ""{
		return nil,ErrMissingArgs
	}

	return nil,nil

}