package data

type Contract struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Channel struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Contracts []Contract `json:"contracts"`
}

type Peer struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Channels []Channel `json:"channels"`
}

type Organization struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Peers []Peer `json:"peers"`
}
