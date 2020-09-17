package component

import (
	"math"
	"math/rand"
)

func NewFaker(start, upper, bottom, volatilityAbs float64) *faker {
	return &faker{
		start:         start,
		upper:         upper,
		bottom:        bottom,
		volatilityAbs: volatilityAbs,

		last: start,
	}
}

type faker struct {
	start         float64
	upper         float64
	bottom        float64
	volatilityAbs float64

	last float64
}

func (f *faker) Next() (bid, ask float64) {
	hv := f.volatilityAbs / 2
	rmin := math.Min(f.bottom, f.last-hv)
	rmax := math.Max(f.upper, f.last+hv)

	f.last = rmin + rand.Float64()*(rmax-rmin)

	bid = f.last - hv
	ask = f.last + hv

	return bid, ask
}
