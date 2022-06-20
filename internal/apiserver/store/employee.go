package store

import (
	"context"

	metav1 "github.com/FuradWho/Cherry/internal/pkg/model/goserver/v1"
)

type EmployeeStore interface {
	Register(ctx context.Context, employee *metav1.Employee) error
}
