package builderpattern

type F16 struct {
	Engine  string
	Wings   string
	Cockpit string
}

type F16Builder struct {
	aircraft *F16
}

func NewF16Builder() *F16Builder {
	return &F16Builder{
		aircraft: &F16{},
	}
}

func (f *F16Builder) BuildEngine() {
	f.aircraft.Engine = "F16 Engine"
}

func (f *F16Builder) BuildWings() {
	f.aircraft.Wings = "F16 Wings"
}

func (f *F16Builder) BuildCockpit() {
	f.aircraft.Cockpit = "F16 Cockpit"
}

func (f *F16Builder) BuildBathrooms() {
	// do nothing
}

func (f *F16Builder) GetAircraft() *F16 {
	return f.aircraft
}
