package go_double

import "testing"

type StubProductRepository struct {
	product Product
}

func (repository StubProductRepository) FindById(id string) *Product {
	return &repository.product
}

func TestFindByIdSholdReturnProduct(t *testing.T) {
	give := "0"
	want := Product{id: "0", ProductName: "SomeProduct"}

	productService := ProductService{}
	productRepository := StubProductRepository{want}

	result, err := productService.FindById(productRepository, give)

	if result.ProductName != want.ProductName {
		t.Errorf("Want '%s' got '%s'", result.ProductName, want.ProductName)
	}

	if err != nil {
		t.Errorf("Error Should Be Nil")
	}

}
