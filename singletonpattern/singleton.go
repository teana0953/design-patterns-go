package singletonpattern

import (
	"log/slog"
	"sync"
)

var aircraftOne *Aircraft
var once sync.Once

type Aircraft struct {
	// some fields
}

func (a *Aircraft) Fly() {
	slog.Info("aircraft is flying")
}

func newAirCraft() *Aircraft {
	return &Aircraft{}
}

func GetAircraft() *Aircraft {
	once.Do(func() {
		aircraftOne = newAirCraft()
	})
	return aircraftOne
}
