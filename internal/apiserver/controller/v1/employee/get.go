package employee

import (
	"github.com/FuradWho/Cherry/internal/pkg/core"
	"github.com/FuradWho/Cherry/internal/pkg/errno"
	metav1 "github.com/FuradWho/Cherry/internal/pkg/model/goserver/v1"
	"github.com/gin-gonic/gin"
)

// Get  get an user by the user identifier.
func (u *EmployeesController) Get(c *gin.Context) {

	var r metav1.Employee

	err := c.ShouldBindJSON(&r)
	if err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	err = u.srv.Employees().Register(c, &r)
	if err != nil {
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
