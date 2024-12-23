package simulators

import (
	"fmt"
)

type Vessel interface {
	FluidMass() float64
	Pressure() float64
}

type flowRateFunc func(Vessel) float64
type updateFunc[K any] func(K, float64, float64) error

/**
* Holds the setup of the simulation.
* The simulation can be run over and over again by calling Simulate().
 */
type DrainSimulator[K Vessel] struct {
	vessel   K
	timeStep float64
	// Takes the vessel and returns what the flow rate out of it is.
	flowRateFunc flowRateFunc
	updateFunc   updateFunc[K]
}

func NewDrainSimulator[K Vessel](vessel K, updateFunc updateFunc[K], timeStep float64, flowRateFunc flowRateFunc) *DrainSimulator[K] {
	return &DrainSimulator[K]{
		timeStep:     timeStep,
		vessel:       vessel,
		flowRateFunc: flowRateFunc,
		updateFunc:   updateFunc,
	}
}

func (d *DrainSimulator[K]) Simulate() []DrainSimulationRecord {
	fmt.Println("Starting drain simulation")
	result := make([]DrainSimulationRecord, 0)
	currentTime := float64(0)
	iters := 0

	for {
		massChange := d.flowRateFunc(d.vessel) * d.timeStep
		if d.vessel.FluidMass()+massChange < 0 {
			fmt.Printf("Stopping drain because massChange=%f and fluidMass=%f", massChange, d.vessel.FluidMass())
			break
		}

		d.updateFunc(d.vessel, d.timeStep, massChange)

		currentTime += d.timeStep
		record := DrainSimulationRecord{
			currentTime,
			d.vessel.FluidMass(),
			d.vessel.Pressure(),
		}

		result = append(result, record)

		iters++
		if iters%50 == 0 {
			fmt.Println(record)
		}
	}

	fmt.Printf("Finished drain simulation. Collected %d records. \n", len(result))

	return result
}

type DrainSimulationRecord struct {
	Time float64
	Mass float64
	// Pressure in Pa.
	Pressure float64
}

func (d DrainSimulationRecord) String() string {
	return fmt.Sprintf("{ Time: %.2f s, Mass: %.2f kg, Pressure: %.2f Pa }", d.Time, d.Mass, d.Pressure)
}
