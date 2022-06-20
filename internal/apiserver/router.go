package apiserver

import (
	"github.com/FuradWho/Cherry/internal/apiserver/controller/v1/employee"
	"github.com/FuradWho/Cherry/internal/apiserver/store/mysql"
	"github.com/FuradWho/Cherry/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

// loadRouter loads the middlewares, routes, handlers.
func loadRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	installMiddleware(g, mw...)
	installController(g)
}

// installMiddleware install Middlewares.
func installMiddleware(g *gin.Engine, mw ...gin.HandlerFunc) {

	g.Use(gin.Recovery())
	//g.Use(middleware.NoCache)
	//g.Use(middleware.Options)
	//g.Use(middleware.Secure)
	g.Use(mw...)

}

// installController install controllers.
func installController(g *gin.Engine) {

	// 404 Handler.
	//g.NoRoute(func(c *gin.Context) {
	//	core.WriteResponse(c, errno.ErrPageNotFound, nil)
	//})

	// /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	storeIns, _ := mysql.GetMySQLFactoryOr()
	employeesController := employee.NewEmployeesController(storeIns)

	v1 := g.Group("/v1")
	{
		employeeV1 := v1.Group("/employees")
		{
			employeeV1.POST("register", employeesController.Register)
			employeeV1.POST("nft_application", employeesController.NFTApplication)
			employeeV1.GET("nft_audit", employeesController.NFTAudit)

			// NFTApplication
		}
	}

}
