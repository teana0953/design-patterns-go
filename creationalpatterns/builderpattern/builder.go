package builderpattern

type AircraftBuilder interface {
	BuildEngine()
	BuildWings()
	BuildCockpit()
	BuildBathrooms()
}
