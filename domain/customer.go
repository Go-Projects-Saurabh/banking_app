package domain

type (
	Customer struct {
		Id          string
		Name        string
		City        string
		Zipcode     string
		DateofBirth string
		Status      string
	}
	//Repository <interface>
	CustomerRepository interface {
		FindAll() ([]Customer, error)
		ById(string) (*Customer, error)
	}
)
