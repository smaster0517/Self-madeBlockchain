package app

import (
	"net/http"

	"github.com/JungBin-Eom/Mini-BlockChain/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()

	a := &AppHandler{
		Handler: r,
		db:      model.NewDBHandler(),
	}

	r.HandleFunc("/peer", a.CreatePeerHandler).Methods("POST")
	r.HandleFunc("/chan", a.CreateChanHandler).Methods("POST")
	r.HandleFunc("/contract", a.CreateContractHandler).Methods("POST")
	r.HandleFunc("/org", a.CreateOrgHandler).Methods("POST")

	r.HandleFunc("/peer", a.GetPeer).Methods("GET")
	r.HandleFunc("/chan", a.GetChan).Methods("GET")
	r.HandleFunc("/contract", a.GetContract).Methods("GET")
	r.HandleFunc("/org", a.GetOrg).Methods("GET")

	r.HandleFunc("/chan/{chan}:[A-z]+}")

	return a
}
