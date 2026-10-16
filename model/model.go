package model

import "sync"

// modelo que deve se seguir
type Url struct {
	Url		string `json:"url"`
	Code	string  `json:"code"`
	Clicks int `json:"clicked"`
}

var (
	Mu sync.Mutex
	Counter int
)

var Link = Url{
		Url: "",
		Code: "",
		Clicks: 0,
}
