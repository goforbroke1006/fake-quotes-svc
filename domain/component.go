package domain

type Faker interface {
	Next() (quote float64)
}

type Emitter interface {
	Emit() Quote
}
