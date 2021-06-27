package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *St) router() http.Handler {
	r := mux.NewRouter()

	mrs := a.mwfRequestSession

	r.HandleFunc("/tokens/{platform:(?:ios|android|web)}/{token:[^/]+}", mrs(a.hTokenCreate)).Methods("POST")
	r.HandleFunc("/tokens/{token:[^/]+}", a.hTokenDestroy).Methods("DELETE")
	r.HandleFunc("/usrs/{id:[0-9]+}", a.hUsrDestroy).Methods("DELETE")
	r.HandleFunc("/send", a.hSend).Methods("POST")

	return a.middleware(r)
}
