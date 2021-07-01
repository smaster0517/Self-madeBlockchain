package app

import (
	"encoding/json"
	"net/http"

	"github.com/JungBin-Eom/Mini-BlockChain/data"
)

func (a *AppHandler) CreateContractHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	a.db.AddContract(name)
}

func (a *AppHandler) GetContract(rw http.ResponseWriter, r *http.Request) {
	list := a.db.GetContract()
	rd.JSON(rw, http.StatusOK, list)
}

func (a *AppHandler) AddFunction(rw http.ResponseWriter, r *http.Request) {
	var req data.FuncRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
	a.db.AddFunction(req)
}
