package model

import "github.com/JungBin-Eom/Mini-BlockChain/data"

type DBHandler interface {
	AddPeer(string)
	AddChannel(string)
	AddContract(string)
	AddOrganization(string)

	GetPeer() []*data.Peer
	GetChan() []*data.Channel
	GetContract() []*data.Contract
	GetOrg() []*data.Organization

	Close()
}

func NewDBHandler() DBHandler {
	return newMemoryHandler()
}
