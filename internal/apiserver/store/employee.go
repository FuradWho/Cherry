package store

import "context"

type EmployeeStore interface {

	// , user *v1.User, opts metav1.CreateOptions

	Create(ctx context.Context) error
}
