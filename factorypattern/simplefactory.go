package factorypattern

type F16SimpleFactory struct{}

type F16A struct{}

func (f *F16A) Fly() string {
	return "F16A"
}

type F16B struct{}

func (f *F16B) Fly() string {
	return "F16B"
}

func (f *F16SimpleFactory) MakeF16(variant string) Aircraft {
	switch variant {
	case "A":
		return &F16A{}
	case "B":
		return &F16B{}
	default:
		return nil
	}
}
