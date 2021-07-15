package data

type Contract struct {
	Id       int                                     `json:"id"`
	Name     string                                  `json:"name"`
	Function []func([]Peer, []Peer) (int, int, bool) `json:"-"`
}

type Channel struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Contracts []Contract `json:"contracts"`
	Peers     []Peer     `json:"peers"`
}

type Peer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Organization struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Peers []Peer `json:"peers"`
}

type JoinRequest struct {
	Channel string `json:"channel"`
	Object  string `json:"object"`
}

type FuncRequest struct {
	Contract string `json:"contract"`
	Function string `json:"function"`
}

type ExcuteRequest struct {
	Channel  string   `json:"channel"`
	Contract string   `json:"contract"`
	Function string   `json:"function"`
	SensorP  []string `json:"sensor_peers"`
	ServiceP []string `json:"service_peers"`
}

type BlockRequest struct {
	Time        string `json:"time"`
	ChannelName string `json:"channel_name"`
	SensorVal   string `json:"sensor_val"`
	ServiceVal  string `json:"service_val"`
}
