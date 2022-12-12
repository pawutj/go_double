package go_double

import "testing"

type SpyProductRepository struct {
	product       Product
	findWasCalled bool
}

func (repository *SpyProductRepository) FindById(id string) *Product {
	repository.findWasCalled = true
	return &repository.product
}

func TestFindByIdSholdReturnProductAndFindByIdShouldBeCalled(t *testing.T) {
	give := "0"
	want := Product{id: "0", ProductName: "SomeProduct"}

	productService := ProductService{}
	productRepository := SpyProductRepository{want, false}

	result, err := productService.FindById(&productRepository, give)

	if result.ProductName != want.ProductName {
		t.Errorf("Want '%s' got '%s'", result.ProductName, want.ProductName)
	}

	if err != nil {
		t.Errorf("Error Should Be Nil")
	}

	if productRepository.findWasCalled != true {
		t.Errorf("findWasCalled Should Be true")
	}

}
