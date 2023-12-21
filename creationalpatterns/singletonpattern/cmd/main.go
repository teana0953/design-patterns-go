package main

import (
	"sync"

	singleton "github.com/teana0953/design-patterns-go/creationalpatterns/singletonpattern"
)

// can use "go run -race ..." to test race condition
func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// not safe
			// notSafeAircraft := singleton.GetNotSafeAircraft()
			// notSafeAircraft.Fly()

			// mutex
			// aircraftMutex := singleton.GetAircraftMutex()
			// aircraftMutex.Fly()

			// check-lock-check
			// aircraftMutexDoubleCheck := singleton.GetAircraftMutexDoubleCheck()
			// aircraftMutexDoubleCheck.Fly()

			// sync.Once
			aircraft := singleton.GetAircraft()
			aircraft.Fly()
		}(wg)
	}
	wg.Wait()
}
