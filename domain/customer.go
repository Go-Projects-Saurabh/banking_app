package domain

import "github.com/saurabhsingh1408/banking_app/errs"

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
		ById(string) (*Customer, *errs.AppError)
	}
)
