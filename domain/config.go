package domain

type Configuration struct {
	Actives []Active
}

type Active struct {
	Code string
	Opts ValueOpts
}

type ValueOpts struct {
	Start         float64
	Upper         float64
	Bottom        float64
	VolatilityAbs float64
}
