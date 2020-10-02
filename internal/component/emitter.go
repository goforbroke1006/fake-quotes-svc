package component

import (
	"math/rand"
	"time"

	"github.com/goforbroke1006/fake-quotes-svc/domain"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/wshub"
)

func NewEmitter(
	active domain.Active,
	hub *wshub.WSHub,
	frequency uint,
) *emitter {
	return &emitter{
		active: active,
		faker: NewFaker(
			active.Opts.Start,
			active.Opts.Upper,
			active.Opts.Lower,
		),
		hub:       hub,
		frequency: frequency,
	}
}

type emitter struct {
	active    domain.Active
	faker     domain.QuoteFaker
	hub       *wshub.WSHub
	frequency uint
}

func (e emitter) Emit() {
	for {
		quote := e.faker.Next()
		q := domain.Quote{
			Code:  e.active.Code,
			Value: quote,
			At:    time.Now().Unix(),
		}
		e.hub.Send(q)

		dur := float64(e.frequency)/2 + rand.Float64()*(float64(e.frequency)/2)

		time.Sleep(time.Duration(dur) * time.Millisecond)
	}
}
