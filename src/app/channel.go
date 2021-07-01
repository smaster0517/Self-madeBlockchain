package app

import (
	"net/http"
)

func (a *AppHandler) CreateChanHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	a.db.AddChannel(name)
}

func (a *AppHandler) GetChan(rw http.ResponseWriter, r *http.Request) {
	list := a.db.GetChan()
	rd.JSON(rw, http.StatusOK, list)
}
