package go_double

import "testing"

type FakeProductRepository struct{}

func (repository *FakeProductRepository) FindById(id string) *Product {

	if id == "0" {
		return &Product{ProductName: "SomeProduct0", id: "0"}
	}

	if id == "1" {
		return &Product{ProductName: "SomeProduct1", id: "1"}
	}

	return &Product{}
}

func TestFindById0SholdReturnProduct0(t *testing.T) {
	give := "0"
	want := Product{ProductName: "SomeProduct0", id: "0"}

	productService := ProductService{}
	productRepository := FakeProductRepository{}

	result, err := productService.FindById(&productRepository, give)

	if result.ProductName != want.ProductName {
		t.Errorf("Want '%s' got '%s'", result.ProductName, want.ProductName)
	}

	if err != nil {
		t.Errorf("Error Should Be Nil")
	}

}

func TestFindById1SholdReturnProduct1(t *testing.T) {
	give := "1"
	want := Product{ProductName: "SomeProduct1", id: "1"}

	productService := ProductService{}
	productRepository := FakeProductRepository{}

	result, err := productService.FindById(&productRepository, give)

	if result.ProductName != want.ProductName {
		t.Errorf("Want '%s' got '%s'", result.ProductName, want.ProductName)
	}

	if err != nil {
		t.Errorf("Error Should Be Nil")
	}

}
