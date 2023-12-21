package main

import (
	"log/slog"

	builder "github.com/teana0953/design-patterns-go/creationalpatterns/builderpattern"
)

func main() {
	F16Builder := builder.NewF16Builder()
	Director := builder.NewDirector(F16Builder)
	Director.Construct(false)

	f16 := F16Builder.GetAircraft()
	slog.Info("f16 is built", slog.Any("f16", f16))

	Boeing747Builder := builder.NewBoeing747Builder()
	Director = builder.NewDirector(Boeing747Builder)
	Director.Construct(true)

	boeing747 := Boeing747Builder.GetAircraft()
	slog.Info("boeing747 is built", slog.Any("boeing747", boeing747))
}
