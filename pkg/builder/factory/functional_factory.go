package factory

// Functional factory is a type of factory that returns functions.
// Later these functions can be used to generate objects.

// This is useful, when you want to create specialized instances of object.

type PowerBand int

const (
	Cost PowerBand = iota
	Performance
	Critical
)

func NewSpotInstanceFactory(band PowerBand) func(string, PowerState) Operator {
	return func(name string, state PowerState) Operator {
		return &machine{name, state, band}
	}
}
