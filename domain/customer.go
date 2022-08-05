package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

//Repository <interface>
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
