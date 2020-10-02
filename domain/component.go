package domain

type QuoteFaker interface {
	Next() (quote float64)
}
