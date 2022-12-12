package go_double

import "testing"

type MockProductRepository struct {
	product       Product
	methodsToCall map[string]bool
}

func (repository *MockProductRepository) FindById(id string) *Product {
	repository.methodsToCall["FindById"] = true
	return &repository.product
}

func (repository *MockProductRepository) ExpectToCall(methodName string) {
	if repository.methodsToCall == nil {
		repository.methodsToCall = make(map[string]bool)
	}
	repository.methodsToCall[methodName] = false
}

func (repository *MockProductRepository) Verify(t *testing.T) {
	for methodName, called := range repository.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindByIdSholdReturnProductUsingMock(t *testing.T) {
	give := "0"
	want := Product{id: "0", ProductName: "SomeProduct"}

	productService := ProductService{}
	productRepository := MockProductRepository{want, nil}

	productRepository.ExpectToCall("FindById")

	result, _ := productService.FindById(&productRepository, give)

	if result.ProductName != want.ProductName {
		t.Errorf("Want '%s' got '%s'", result.ProductName, want.ProductName)
	}

	productRepository.Verify(t)

}
