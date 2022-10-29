package rest

// @Router  /send [post]
// @Tags    send
// @Param   body body entities.SendReqSt false "body"
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
// func (o *St) hSend(c *gin.Context) {
// 	reqObj := &entities.SendReqSt{}
// 	if !dopHttps.BindJSON(c, reqObj) {
// 		return
// 	}
//
// 	err := o.core.Send(reqObj)
// 	if dopHttps.Error(c, err) {
// 		return
// 	}
// }
