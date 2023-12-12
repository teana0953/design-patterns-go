package builderpattern

// product of the builder pattern
type Boeing747 struct {
	Engine    string
	Wings     string
	Cockpit   string
	Bathrooms string
}

// concrete builder of the builder pattern
type Boeing747Builder struct {
	aircraft *Boeing747
}

func NewBoeing747Builder() *Boeing747Builder {
	return &Boeing747Builder{
		aircraft: &Boeing747{},
	}
}

func (b *Boeing747Builder) BuildEngine() {
	b.aircraft.Engine = "Boeing747 Engine"
}

func (b *Boeing747Builder) BuildWings() {
	b.aircraft.Wings = "Boeing747 Wings"
}

func (b *Boeing747Builder) BuildCockpit() {
	b.aircraft.Cockpit = "Boeing747 Cockpit"
}

func (b *Boeing747Builder) BuildBathrooms() {
	b.aircraft.Bathrooms = "Boeing747 Bathrooms"
}

func (b *Boeing747Builder) GetAircraft() *Boeing747 {
	return b.aircraft
}
