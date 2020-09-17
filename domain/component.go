package domain

type Faker interface {
	Next() (bid, ask float64)
}

type Emitter interface {
	Emit() Quote
}
