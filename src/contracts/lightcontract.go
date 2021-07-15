package contracts

import (
	"github.com/JungBin-Eom/Mini-BlockChain/data"
)

func LightAdjust(sensors []data.Peer, services []data.Peer) (int, int, bool) {
	sum := 0
	for _, v := range sensors {
		sum += v.Value
	}
	avg := sum / len(sensors)

	var serviceTarget string
	if avg >= 20 {
		serviceTarget = "lightDown"
	} else if avg <= 10 {
		serviceTarget = "lightUp"
	}

	if serviceTarget == "" {
		return avg, 0, false
	}

	var targetNum int
	for i, v := range services {
		if v.Name == serviceTarget {
			targetNum = i
			break
		}
	}

	return avg, targetNum, true
}
