package day11

import (
	"fmt"
	"strings"
)

type device struct {
	label   string
	outputs []*device
}

type serverRoom struct {
	devices map[string]*device
}

func parseInput(lines []string) serverRoom {
	devices := make(map[string]*device)

	for _, line := range lines {
		labels := strings.Fields(line)

		thisDevice := labels[0][:len(labels[0])-1]

		if _, present := devices[thisDevice]; present { // We already know about this machine - fill in the outputs
			for _, output := range labels[1:] {
				if _, present := devices[output]; present { // The output is also known, we can add the pointer to the slice
					localOutputs := devices[thisDevice].outputs
					localOutputs = append(localOutputs, devices[output])
					devices[thisDevice].outputs = localOutputs
				} else { // The output is not yet know, add it and point to it
					newDevice := &device{
						label:   output,
						outputs: []*device{},
					}

					devices[output] = newDevice

					localOutputs := devices[thisDevice].outputs
					localOutputs = append(localOutputs, newDevice)
					devices[thisDevice].outputs = localOutputs
				}
			}
		} else { // We don't know about this device yet - create it
			newDevice := &device{
				label:   thisDevice,
				outputs: []*device{},
			}

			for _, output := range labels[1:] {
				if _, present := devices[output]; present { // The output is already known, we can add the pointer to the slice
					localOutputs := newDevice.outputs
					localOutputs = append(localOutputs, devices[output])
					newDevice.outputs = localOutputs
				} else { // The output is not yet know, add it and point to it
					newOutput := &device{
						label:   output,
						outputs: []*device{},
					}

					devices[output] = newOutput

					localOutputs := newDevice.outputs
					localOutputs = append(localOutputs, newOutput)
					newDevice.outputs = localOutputs
				}
			}

			devices[thisDevice] = newDevice
		}
	}

	return serverRoom{
		devices: devices,
	}
}

func (sr serverRoom) countPaths(start, stop string) int {
	var out int

	queue := make([]*device, 0)

	present := sr.devices[start]

	for _, output := range present.outputs {
		// Add these next steps to the queue
		queue = append(queue, output)
	}

	for len(queue) != 0 {
		present = queue[0]
		queue = queue[1:]

		if present.label == stop {
			out++
		} else {
			for _, output := range present.outputs {
				queue = append(queue, output)
			}
		}
	}

	return out
}

func (sr serverRoom) routesPassingThrough(start, stop string) int {
	var out int

	type state struct {
		device *device
		dac    bool
		fft    bool
	}

	queue := make([]state, 0)

	present := state{
		device: sr.devices[start],
	}

	for _, output := range present.device.outputs {
		// Add these next steps to the queue
		queue = append(queue, state{
			device: output,
		})
	}

	for len(queue) != 0 {
		present = queue[0]
		queue = queue[1:]

		if present.device.label == stop {
			if present.dac && present.fft {
				fmt.Println("Adding one!")
				out++
			}
		} else {
			for _, output := range present.device.outputs {
				dac := present.dac || present.device.label == "dac"
				fft := present.fft || present.device.label == "fft"
				queue = append(queue, state{
					device: output,
					dac:    dac,
					fft:    fft,
				})
			}
		}
	}

	return out
}
