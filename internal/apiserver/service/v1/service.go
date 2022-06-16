package v1

import "github.com/FuradWho/Cherry/internal/apiserver/store"

// Service defines functions used to return resource interface.
type Service interface {
	Employees() EmployeeSrv
}

type service struct {
	store store.Factory
}

func (s *service) Employees() EmployeeSrv {
	return newEmployees(s)
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
