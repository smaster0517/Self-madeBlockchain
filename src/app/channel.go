package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	block := a.db.ExecuteFunction(req)
	if block.ChannelName == "" {
		rd.Text(rw, http.StatusOK, "No Block Create")
		return
	}
	bbytes, err := json.Marshal(block)
	if err != nil {
		http.Error(rw, "Unable to marshal block struct", http.StatusInternalServerError)
	}
	buff := bytes.NewBuffer(bbytes)
	fmt.Println("block request")
	res, err := http.Post("http://3.35.174.241:8000/block", "application/json", buff)
	if err != nil {
		http.Error(rw, "Unable to request creating block", http.StatusBadRequest)
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err == nil {
		rd.Text(rw, http.StatusOK, string(resBody))
	} else {
		rd.Text(rw, http.StatusInternalServerError, "Unable to read body")
	}
}

func (a *AppHandler) GetBlock(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channel, _ := vars["channel"]
	req, err := http.NewRequest("GET", "http://3.35.174.241:8000/block?channel_name="+channel, nil)
	if err != nil {
		http.Error(rw, "Unable to get block", http.StatusBadRequest)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(rw, "Unable to do request", http.StatusInternalServerError)
	}
	defer res.Body.Close()

	fmt.Println(res.Body)

	resBody, err := ioutil.ReadAll(res.Body)
	if err == nil {
		rd.Text(rw, http.StatusOK, string(resBody))
	} else {
		rd.Text(rw, http.StatusInternalServerError, "Unable to get block")
	}
}
