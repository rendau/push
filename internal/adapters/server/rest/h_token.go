package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
	"github.com/rendau/push/internal/domain/entities"
)

// @Router  /token [post]
// @Tags    token
// @Param   body body     entities.TokenCUSt false "body"
// @Success 200  {object} dopTypes.CreateRep{id=string}
// @Failure 400  {object} dopTypes.ErrRep
func (o *St) hTokenCreate(c *gin.Context) {
	reqObj := &entities.TokenCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.core.Token.Create(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router  /token/:value [delete]
// @Tags    token
// @Param   value path string true "id"
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hTokenDelete(c *gin.Context) {
	value := c.Param("value")

	dopHttps.Error(c, o.core.Token.Delete(o.getRequestContext(c), value))
}
