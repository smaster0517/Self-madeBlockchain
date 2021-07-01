package model

import (
	"github.com/JungBin-Eom/Mini-BlockChain/data"
)

type memoryHandler struct {
	peerMap         map[int]*data.Peer
	channelMap      map[int]*data.Channel
	contractMap     map[int]*data.Contract
	organizationMap map[int]*data.Organization
}

func (m *memoryHandler) AddPeer(name string) {
	var p data.Peer
	p.Name = name
	p.Id = len(m.peerMap) + 1
	p.Channels = make([]data.Channel, 0)

	m.peerMap[p.Id] = &p
}

func (m *memoryHandler) AddChannel(name string) {
	var p data.Channel
	p.Name = name
	p.Id = len(m.peerMap) + 1
	p.Contracts = make([]data.Contract, 0)

	m.channelMap[p.Id] = &p
}

func (m *memoryHandler) AddContract(name string) {
	var p data.Contract
	p.Name = name
	p.Id = len(m.peerMap) + 1

	m.contractMap[p.Id] = &p
}

func (m *memoryHandler) AddOrganization(name string) {
	var p data.Organization
	p.Name = name
	p.Id = len(m.peerMap) + 1
	p.Peers = make([]data.Peer, 0)

	m.organizationMap[p.Id] = &p
}

func (m *memoryHandler) GetPeer() []*data.Peer {
	list := []*data.Peer{}
	for _, v := range m.peerMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) GetChan() []*data.Channel {
	list := []*data.Channel{}
	for _, v := range m.channelMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) GetContract() []*data.Contract {
	list := []*data.Contract{}
	for _, v := range m.contractMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) GetOrg() []*data.Organization {
	list := []*data.Organization{}
	for _, v := range m.organizationMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) Close() {

}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.peerMap = make(map[int]*data.Peer)
	m.channelMap = make(map[int]*data.Channel)
	m.contractMap = make(map[int]*data.Contract)
	m.organizationMap = make(map[int]*data.Organization)
	return m
}
