package singletonpattern

var aircraftNotSafe *Aircraft

func GetNotSafeAircraft() *Aircraft {
	if aircraftNotSafe == nil {
		aircraftNotSafe = newAirCraft()
	}
	return aircraftNotSafe
}
