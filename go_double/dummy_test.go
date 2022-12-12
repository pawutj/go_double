package go_double

import "testing"

type DummyProductRepository struct{}

func (repository *DummyProductRepository) FindById(id string) *Product {
	return &Product{}
}

func TestFindByIdSholdReturnErrorWhenIdNil(t *testing.T) {
	productService := ProductService{}
	productRepository := DummyProductRepository{}
	want := ErrMissingArgs

	_, got := productService.FindById(&productRepository, "")

	if got != ErrMissingArgs {
		t.Errorf("Want '%s' got '%s'", want, got)
	}

}
