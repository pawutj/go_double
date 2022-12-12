package go_double

type DB struct {}

type ProductRepository interface{
	FindById(db DB , id string) *Product
}
