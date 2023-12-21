package abstractfactorypattern

type Aircraft interface {
	Fly() string
}

type AircraftFactory interface {
	MakeAircraft() Aircraft
}

type F16AFactory struct{}

func (f *F16AFactory) MakeAircraft() Aircraft {
	return &F16A{}
}

type F16BFactory struct{}

func (f *F16BFactory) MakeAircraft() Aircraft {
	return &F16B{}
}
