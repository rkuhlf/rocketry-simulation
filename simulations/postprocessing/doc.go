/**
* Post-processing does depend directly on the simulator, but that's because there is an inherent coupling there. If the simulation changes, we will always have to change the post-processing, and we never want to reuse these post-processing functions for anything other than the simulations they are intended for.
* It is not directly included within the simulator package because not everyone who uses the simulators package will (in theory) want to postprocess it the way we've done here. So we don't always include the postprocessing with the simulator.
 */
package postprocessing
