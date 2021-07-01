package app

import (
	"encoding/json"
	"net/http"

	"github.com/JungBin-Eom/Mini-BlockChain/data"
	"github.com/gorilla/mux"
)

func (a *AppHandler) CreateChanHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	a.db.AddChannel(name)
}

func (a *AppHandler) GetChan(rw http.ResponseWriter, r *http.Request) {
	list := a.db.GetChan()
	rd.JSON(rw, http.StatusOK, list)
}

func (a *AppHandler) JoinPeer(rw http.ResponseWriter, r *http.Request) {
	var join data.JoinRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&join)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
	if a.db.JoinPeer(join) {
		rd.Text(rw, http.StatusOK, "success")
	} else {
		rd.Text(rw, http.StatusInternalServerError, "fail")
	}
}

func (a *AppHandler) JoinContract(rw http.ResponseWriter, r *http.Request) {
	var join data.JoinRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&join)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
	if a.db.JoinContract(join) {
		rd.Text(rw, http.StatusOK, "success")
	} else {
		rd.Text(rw, http.StatusInternalServerError, "fail")
	}
}

func (a *AppHandler) ExecuteFunction(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req data.ExcuteRequest
	req.Channel = vars["channel"]
	req.Contract = vars["contract"]
	req.Function = vars["function"]
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	if a.db.ExecuteFunction(req) {
		rd.Text(rw, http.StatusOK, "success")
	} else {
		rd.Text(rw, http.StatusInternalServerError, "fail")
	}
}
