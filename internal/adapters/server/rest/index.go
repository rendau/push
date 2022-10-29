package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rendau/dop/adapters/logger"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/push/internal/domain/core"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

type St struct {
	lg   logger.Lite
	core *core.St
}

func GetHandler(lg logger.Lite, core *core.St, withCors bool) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares

	r.Use(dopHttps.MwRecovery(lg, nil))
	if withCors {
		r.Use(dopHttps.MwCors())
	}

	// handlers

	// doc
	r.GET("/doc/*any", ginSwag.WrapHandler(swagFiles.Handler, func(c *ginSwag.Config) {
		c.DefaultModelsExpandDepth = 0
		c.DocExpansion = "none"
	}))

	s := &St{lg: lg, core: core}

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// token
	r.POST("/token", s.hTokenCreate)
	r.DELETE("/token/:value", s.hTokenDelete)

	// usr
	r.DELETE("/usr/:id/token", s.hUsrTokenDelete)

	// main
	// r.POST("/send", s.hSend)

	return r
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return o.core.Session.SetToContextByToken(context.Background(), dopHttps.GetAuthToken(c))
}
