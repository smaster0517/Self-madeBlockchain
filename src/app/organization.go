package app

import (
	"net/http"
)

func (a *AppHandler) CreateOrgHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	a.db.AddOrganization(name)
}

func (a *AppHandler) GetOrg(rw http.ResponseWriter, r *http.Request) {
	list := a.db.GetOrg()
	rd.JSON(rw, http.StatusOK, list)
}
