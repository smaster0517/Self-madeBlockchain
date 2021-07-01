package contracts

import (
	"github.com/JungBin-Eom/Mini-BlockChain/data"
)

func LightUp(device *data.Peer, service *data.Peer) (int, int) {
	if device.Value < 10 {
		service.Value += 1
	} else if device.Value > 20 {
		service.Value -= 1
	}
	return device.Value, service.Value
}
