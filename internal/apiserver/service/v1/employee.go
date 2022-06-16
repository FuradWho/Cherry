package v1

import (
	"context"
	"fmt"
	"github.com/FuradWho/Cherry/internal/apiserver/store"
)

type EmployeeSrv interface {

	// , user *v1.User, opts metav1.CreateOptions

	Create(ctx context.Context) error
}

type employeeService struct {
	store store.Factory
}

func newEmployees(srv *service) *employeeService {
	return &employeeService{store: srv.store}
}

func (s *employeeService) Create(ctx context.Context) error {
	fmt.Println(1)
	return nil
}
