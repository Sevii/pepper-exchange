package main

import ()

//FillBus distributes fill output from the orderbook to consumers
type FillBus struct {
	subscribers []chan Fill
}

func NewFillBus() *FillBus {
	return &FillBus{}
}

func (d *FillBus) Run(in <-chan Fill) {
	for fill := range in {
		//WAL

		for _, sub := range d.subscribers {

			sub <- fill
		}
	}
}

func (bus *FillBus) subscribe(out chan Fill) {
	//Possible race, luckily subscribers get best effort guarentees
	bus.subscribers = append(bus.subscribers, out)
}
