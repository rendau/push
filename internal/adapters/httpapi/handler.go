package httpapi

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rendau/push/internal/constants"
	"github.com/rendau/push/internal/domain/entities"
	"github.com/rendau/push/internal/errs"
)

func (a *St) hTokenCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var platformId int

	args := mux.Vars(r)

	token := args["token"]
	platform := args["platform"]

	switch platform {
	case "ios":
		platformId = constants.PlatformIOS
	case "android":
		platformId = constants.PlatformAndroid
	case "web":
		platformId = constants.PlatformWeb
	default:
		a.lg.Errorw("Not correct platform", errs.NotCorrectPlatform)
		return
	}

	st := &entities.TokenCreateSt{
		Token:      token,
		PlatformId: platformId,
	}

	err = a.cr.TokenCreate(r.Context(), st)
	if err != nil {
		a.uHandleError(err, w)
		return
	}
}

func (a *St) hTokenDestroy(w http.ResponseWriter, r *http.Request) {
	var err error

	args := mux.Vars(r)
	token := args["token"]

	err = a.cr.TokenDestroy(token)
	if err != nil {
		a.uHandleError(err, w)
		return
	}
}

func (a *St) hSend(w http.ResponseWriter, r *http.Request) {
	reqObj := &entities.SendReqSt{}
	if !a.uParseRequestJSON(w, r, reqObj) {
		return
	}

	a.lg.Infow("hSend", "obj", *reqObj)

	err := a.cr.Send(reqObj)
	if err != nil {
		a.uHandleError(err, w)
		return
	}
}

func (a *St) hUsrDestroy(w http.ResponseWriter, r *http.Request) {
	var err error

	args := mux.Vars(r)
	id, _ := strconv.ParseInt(args["id"], 10, 64)

	err = a.cr.UsrDestroy(id)
	if err != nil {
		a.uHandleError(err, w)
		return
	}
}
