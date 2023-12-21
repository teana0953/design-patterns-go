package main

import (
	factory "github.com/teana0953/design-patterns-go/abstractfactorypattern"
)

func main() {
	// simple factory
	f16Factory := &factory.F16SimpleFactory{}
	f16A := f16Factory.MakeF16("A")
	f16A.Fly()
	f16B := f16Factory.MakeF16("B")
	f16B.Fly()

	// factory method
	f16AFactory := &factory.F16AFactory{}
	f16A = f16AFactory.MakeAircraft()
	f16A.Fly()
	f16BFactory := &factory.F16BFactory{}
	f16B = f16BFactory.MakeAircraft()
	f16B.Fly()
}
