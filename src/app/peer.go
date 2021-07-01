package app

import (
	"net/http"
)

func (a *AppHandler) CreatePeerHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	a.db.AddPeer(name)
}

func (a *AppHandler) GetPeer(rw http.ResponseWriter, r *http.Request) {
	list := a.db.GetPeer()
	rd.JSON(rw, http.StatusOK, list)
}
