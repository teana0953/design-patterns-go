package singletonpattern

import "sync"

var lock = &sync.Mutex{}
var aircraftMutex *Aircraft

func GetAircraftMutex() *Aircraft {
	lock.Lock()
	defer lock.Unlock()
	if aircraftMutex == nil {
		aircraftMutex = newAirCraft()
	}
	return aircraftMutex
}

// still have race condition since golang compiler and runtime reorder instructions
func GetAircraftMutexDoubleCheck() *Aircraft {
	// check
	if aircraftMutex == nil {
		// lock
		lock.Lock()
		defer lock.Unlock()

		// check
		if aircraftMutex == nil {
			aircraftMutex = newAirCraft()
		}
	}
	return aircraftMutex
}
