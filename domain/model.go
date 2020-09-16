package domain

type Quote struct {
	Code string  `json:"code"`
	Bid  float64 `json:"bid"`
	Ask  float64 `json:"ask"`
	At   int64   `json:"at"`
}
