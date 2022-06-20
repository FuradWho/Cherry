package v1

import (
	"context"
	"github.com/FuradWho/Cherry/internal/apiserver/store"
	"github.com/FuradWho/Cherry/internal/pkg/errno"
	metav1 "github.com/FuradWho/Cherry/internal/pkg/model/goserver/v1"
	"regexp"
)

type EmployeeSrv interface {

	// , user *v1.User, opts metav1.CreateOptions

	Register(ctx context.Context, employee *metav1.Employee) error
}

type employeeService struct {
	store store.Factory
}

func newEmployees(srv *service) *employeeService {
	return &employeeService{store: srv.store}
}

func (s *employeeService) Register(ctx context.Context, employee *metav1.Employee) error {
	err := s.store.Employees().Register(ctx, employee)
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}
	return nil
}
