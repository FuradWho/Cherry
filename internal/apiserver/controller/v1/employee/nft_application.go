package employee

import (
	"github.com/FuradWho/Cherry/internal/pkg/core"
	"github.com/FuradWho/Cherry/internal/pkg/errno"
	"github.com/gin-gonic/gin"
)

type NFTRequest struct {
	WalletAddress string `json:"wallet_address"`
	Password      string `json:"password"`
}

func (u *EmployeesController) NFTApplication(c *gin.Context) {

	var r NFTRequest

	err := c.ShouldBindJSON(&r)
	if err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	nftAddress, err := u.srv.Employees().NFTApplication(c, r.WalletAddress, r.Password)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{
		"status":     "ok",
		"nftAddress": nftAddress,
	})
}
