package domain

type Quote struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
	At    int64   `json:"at"`
}
