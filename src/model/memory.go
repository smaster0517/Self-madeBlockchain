package model

import (
	"reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/JungBin-Eom/Mini-BlockChain/contracts"
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
	p.Value = 0

	m.peerMap[p.Id] = &p
}

func (m *memoryHandler) AddChannel(name string) {
	var c data.Channel
	c.Name = name
	c.Id = len(m.channelMap) + 1
	c.Contracts = make([]data.Contract, 0)
	c.Peers = make([]data.Peer, 0)

	m.channelMap[c.Id] = &c
}

func (m *memoryHandler) AddContract(name string) {
	var c data.Contract
	c.Name = name
	c.Id = len(m.contractMap) + 1

	m.contractMap[c.Id] = &c
}

func (m *memoryHandler) AddOrganization(name string) {
	var o data.Organization
	o.Name = name
	o.Id = len(m.organizationMap) + 1
	o.Peers = make([]data.Peer, 0)

	m.organizationMap[o.Id] = &o
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

func (m *memoryHandler) JoinPeer(join data.JoinRequest) bool {
	for _, v := range m.channelMap {
		if v.Name == join.Channel {
			for _, vv := range m.peerMap {
				if vv.Name == join.Object {
					v.Peers = append(v.Peers, *vv)
					return true
				}
			}
		}
	}
	return false
}

func (m *memoryHandler) JoinContract(join data.JoinRequest) bool {
	for _, v := range m.channelMap {
		if v.Name == join.Channel {
			for _, vv := range m.contractMap {
				if vv.Name == join.Object {
					v.Contracts = append(v.Contracts, *vv)
					return true
				}
			}
		}
	}
	return false
}

func (m *memoryHandler) AddFunction(req data.FuncRequest) bool {
	if req.Function == "LightAdjust" {
		for i, v := range m.contractMap {
			if v.Name == req.Contract {
				m.contractMap[i].Function = append(m.contractMap[i].Function, contracts.LightAdjust)
				return true
			}
		}
	}
	return false
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (m *memoryHandler) ExecuteFunction(req data.ExcuteRequest) *data.BlockRequest {
	var sensorNum, serviceNum, chanNum int
	for i, v := range m.peerMap {
		if v.Name == req.Peer1 {
			sensorNum = i
		} else if v.Name == req.Peer2 {
			serviceNum = i
		}
	}

	for i, v := range m.channelMap {
		if v.Name == req.Channel {
			chanNum = i
			break
		}
	}
	var val1, val2 int
	if req.Function == "LightAdjust" {
		val1, val2 = contracts.LightAdjust(m.peerMap[sensorNum], m.peerMap[serviceNum])
	}

	m.peerMap[sensorNum].Value = val1
	m.peerMap[serviceNum].Value = val2

	for i, v := range m.channelMap[chanNum].Peers {
		if v.Name == m.peerMap[sensorNum].Name {
			m.channelMap[chanNum].Peers[i].Value = val1
		} else if v.Name == m.peerMap[serviceNum].Name {
			m.channelMap[chanNum].Peers[i].Value = val2
		}
	}

	blockReq := &data.BlockRequest{}
	blockReq.Time = time.Now().Format("2021-07-07")
	blockReq.ChannelName = m.channelMap[chanNum].Name
	blockReq.SensorVal = strconv.Itoa(m.peerMap[sensorNum].Value)
	blockReq.ServiceVal = strconv.Itoa(m.peerMap[serviceNum].Value)

	return blockReq
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
