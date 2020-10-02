package component

import (
	"math/rand"
)

func NewFaker(start, upper, lower float64) *faker {
	return &faker{
		start: start,
		upper: upper,
		lower: lower,

		last: start,
	}
}

type faker struct {
	start float64
	upper float64
	lower float64

	last float64
}

func (f *faker) Next() (quote float64) {
	f.last = f.lower + rand.Float64()*(f.upper-f.lower)

	return f.last
}
