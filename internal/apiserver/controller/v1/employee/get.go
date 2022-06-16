package employee

import (
	"github.com/FuradWho/Cherry/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

// Get  get an user by the user identifier.
func (u *EmployeesController) Get(c *gin.Context) {

	err := u.srv.Employees().Create(c)
	if err != nil {
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
