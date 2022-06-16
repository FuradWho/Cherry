package employee

import (
	srvv1 "github.com/FuradWho/Cherry/internal/apiserver/service/v1"
	"github.com/FuradWho/Cherry/internal/apiserver/store"
)

// EmployeesController create a user handler used to handle request for user resource.
type EmployeesController struct {
	srv srvv1.Service
}

// NewEmployeesController creates a user handler.
func NewEmployeesController(store store.Factory) *EmployeesController {
	return &EmployeesController{
		srv: srvv1.NewService(store),
	}
}
