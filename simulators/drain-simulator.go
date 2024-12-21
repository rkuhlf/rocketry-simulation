package simulators

type Vessel interface {
	UpdateState(timeStep float64, massChange float64) error
	FluidMass() float64
}

type flowRateFunc func(Vessel) float64

/**
* Holds the setup of the simulation.
* The simulation can be run over and over again by calling Simulate().
 */
type drainSimulator struct {
	vessel   Vessel
	timeStep float64
	// Takes the vessel and returns what the flow rate out of it is.
	flowRateFunc flowRateFunc
}

func DrainSimulator(vessel Vessel, timeStep float64, flowRateFunc flowRateFunc) *drainSimulator {
	return &drainSimulator{
		timeStep:     timeStep,
		vessel:       vessel,
		flowRateFunc: flowRateFunc,
	}
}

func (d *drainSimulator) Simulate() []DrainSimulationRecord {
	result := make([]DrainSimulationRecord, 0)
	currentTime := float64(0)

	for {
		massChange := d.flowRateFunc(d.vessel) * d.timeStep
		if massChange > d.vessel.FluidMass() {
			break
		}

		d.vessel.UpdateState(d.timeStep, massChange)

		currentTime += d.timeStep
		result = append(result, DrainSimulationRecord{
			d.vessel.FluidMass(),
			currentTime,
		})
	}

	return result
}

type DrainSimulationRecord struct {
	time float64
	mass float64
}
