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
	// var sensorNums, serviceNums []int
	var sensorPeers, servicePeers []data.Peer
	var chanNum int
	var flag bool
	for _, v := range m.peerMap {
		flag = false
		for _, vv := range req.SensorP {
			if v.Name == vv {
				sensorPeers = append(sensorPeers, *v)
				flag = true
				break
			}
		}
		if flag == false {
			for _, vv := range req.ServiceP {
				if v.Name == vv {
					servicePeers = append(servicePeers, *v)
					flag = true
					break
				}
			}
		}
	}

	for i, v := range m.channelMap {
		if v.Name == req.Channel {
			chanNum = i
			break
		}
	}

	/*
		Contract의 종류
		1. Sensor 1 : 1 Service Mapping 하나의 센서는 하나의 서비스만 영향
		2. Sensor 1 : N Service Mapping 하나의 센서가 여럿의 서비스에 영향
		3. Sensor N : 1 Service Mapping 여럿의 센서가 하나의 서비스에 영향
		4. Sensor N : N Service Mapping 여럿의 센서가 여럿의 서비스에 영향

		센서 피어의 값은 일단 평균으로 계산하여 모든 서비스 피어에게 전달하는 것으로
		센서 피어의 평균 값에 따라 수행해야 하는 서비스가 다름
		평균 값을 구하고 서비스 결정한 후 서비스 피어 목록에서 일치하는 것을 찾자
		서비스에 영향을 준다 -> 서비스를 수행한다
	*/

	var sensorAvg, serviceIdx int
	var excuted bool
	if req.Function == "LightAdjust" {
		// 결국 Sensor 피어의 값들을 처리해서 서비스에 전달할거면
		// Service 피어들은 컨트랙트에 포함될 필요가 있는가?
		sensorAvg, serviceIdx, excuted = contracts.LightAdjust(sensorPeers, servicePeers)
	}

	var peerIdx int
	for i, v := range m.peerMap {
		if v.Name == servicePeers[serviceIdx].Name {
			m.peerMap[i].Value += 1
			peerIdx = i
			break
		}
	}

	var channelNum, channelPeerNum int
	for i, v := range m.channelMap {
		if v.Name == req.Channel {
			for j, vv := range m.channelMap[i].Peers {
				if vv.Name == servicePeers[serviceIdx].Name {
					channelNum = i
					channelPeerNum = j
				}
			}
		}
	}

	m.channelMap[channelNum].Peers[channelPeerNum].Value = m.peerMap[peerIdx].Value

	if excuted == true {
		blockReq := &data.BlockRequest{}
		blockReq.Time = time.Now().String()
		blockReq.ChannelName = m.channelMap[chanNum].Name
		blockReq.SensorVal = strconv.Itoa(sensorAvg)
		blockReq.ServiceVal = m.peerMap[peerIdx].Name + " " + strconv.Itoa(m.peerMap[peerIdx].Value)
		// Service는 값을 그냥 실행하는 것 뿐인데 서비스의 값을 주어야하는가?
		// 컨트랙트 실행 결과로 실행된 서비스가 무엇인지만 알면 되는것 아닌가?
		// 들어갈거면 서비스 피어 이름이랑 실제 장치의 값이 들어가야 할 것
		return blockReq
	} else {
		return &data.BlockRequest{}
	}
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
