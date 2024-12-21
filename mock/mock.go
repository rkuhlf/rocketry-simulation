package mock

// MockVessel is a manual mock implementation of the Vessel interface
type MockVessel struct {
	UpdateStateFunc func(timeStep float64, massChange float64) error
	FluidMassFunc   func() float64
}

// Mock implementation of UpdateState
func (m *MockVessel) UpdateState(timeStep float64, massChange float64) error {
	if m.UpdateStateFunc != nil {
		return m.UpdateStateFunc(timeStep, massChange)
	}
	return nil // default no-op behavior
}

// Mock implementation of FluidMass
func (m *MockVessel) FluidMass() float64 {
	if m.FluidMassFunc != nil {
		return m.FluidMassFunc()
	}
	return 0 // default return value
}
