package httpapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rendau/push/internal/domain/entities"
	"github.com/rs/cors"
)

var (
	usrAuthHttpClient = http.Client{Timeout: 5 * time.Second}
)

func (a *St) middleware(h http.Handler) http.Handler {
	h = cors.AllowAll().Handler(h)
	h = a.mwRecovery(h)
	return h
}

func (a *St) mwRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cancelCtx, cancel := context.WithCancel(r.Context())
		r = r.WithContext(cancelCtx)
		defer func() {
			if err := recover(); err != nil {
				cancel()
				w.WriteHeader(http.StatusInternalServerError)
				a.lg.Errorw(
					"Panic in http handler",
					err,
					"method", r.Method,
					"path", r.URL,
				)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func (a *St) mwfRequestSession(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ses := &entities.Session{}

		authReq, err := http.NewRequest("GET", a.usrAuthUrl, nil)
		if err == nil {
			for k, vs := range r.Header {
				for _, v := range vs {
					authReq.Header.Add(k, v)
				}
			}

			authReq.URL.RawQuery = r.URL.Query().Encode()

			if resp, err := usrAuthHttpClient.Do(authReq); err == nil {
				defer resp.Body.Close()
				if resp.StatusCode >= 200 || resp.StatusCode < 300 {
					if repBytes, err := ioutil.ReadAll(resp.Body); err == nil {
						var repObj AuthRepSt
						if err = json.Unmarshal(repBytes, &repObj); err == nil {
							ses.Id = repObj.Id
						} else {
							a.lg.Errorw("Bad json from auth response", err, "body", string(repBytes))
						}
					} else {
						a.lg.Errorw("Fail to read auth response body", err)
					}
				} else {
					a.lg.Errorw("Bad status from auth response", nil, "status_code", resp.StatusCode)
				}
			} else {
				a.lg.Errorw("Fail to sent usr-auth request", err)
			}
		} else {
			a.lg.Errorw("Fail to create http-request for usr auth", err)
		}

		r = r.WithContext(a.cr.ContextWithSession(r.Context(), ses))

		hf(w, r)
	}
}
