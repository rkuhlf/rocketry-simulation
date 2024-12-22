package simulators

func ConstantFlowRate(c float64) func(Vessel) float64 {
	return func(_ Vessel) float64 {
		return c
	}
}
