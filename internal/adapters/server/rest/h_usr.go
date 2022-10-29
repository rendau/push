package rest

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
)

// @Router  /usr/:id/token [delete]
// @Tags    usr
// @Param   value path int true "id"
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hUsrTokenDelete(c *gin.Context) {
	usrId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	dopHttps.Error(c, o.core.Usr.TokenDestroy(o.getRequestContext(c), usrId))
}
