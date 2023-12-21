package builderpattern

type Director struct {
	builder AircraftBuilder
}

func NewDirector(builder AircraftBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct(hasPassenger bool) {
	d.builder.BuildCockpit()
	d.builder.BuildEngine()
	d.builder.BuildWings()

	if hasPassenger {
		d.builder.BuildBathrooms()
	}
}
