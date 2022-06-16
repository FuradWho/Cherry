package apiserver

import (
	"github.com/gin-gonic/gin"
)

func NewGoServer() {

	// create the gin engine
	g := gin.New()

	loadRouter(g)

	err := g.Run()
	if err != nil {
		return
	}
}
