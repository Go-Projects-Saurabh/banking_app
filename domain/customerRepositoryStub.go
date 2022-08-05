package domain

// Adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Saurabh", City: "Noida", Zipcode: "201301", DateofBirth: "10/07/1984", Status: "1"},
		{Id: "1002", Name: "Kavyansh", City: "Noida", Zipcode: "201301", DateofBirth: "14/08/2012", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
