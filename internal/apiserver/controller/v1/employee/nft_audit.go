package employee

import (
	"github.com/FuradWho/Cherry/internal/pkg/core"
	"github.com/FuradWho/Cherry/internal/pkg/errno"
	"github.com/gin-gonic/gin"
)

func (u *EmployeesController) NFTAudit(c *gin.Context) {

	addr := c.Query("nft_address")
	if addr == "" {
		core.WriteResponse(c, errno.ErrValidation, nil)
	}

	audit, err := u.srv.Employees().NFTAudit(c, addr)
	if err != nil || audit != "1" {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{
		"status":     audit,
		"nftAddress": addr,
	})
}
