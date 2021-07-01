package app

import (
	"net/http"
)

func (a *AppHandler) CreateContractHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	a.db.AddContract(name)
}

func (a *AppHandler) GetContract(rw http.ResponseWriter, r *http.Request) {
	list := a.db.GetContract()
	rd.JSON(rw, http.StatusOK, list)
}
