package main

import (
	"log/slog"

	prototype "github.com/teana0953/design-patterns-go/prototypepattern"
)

func main() {
	original := prototype.NewConcretePrototype(100)
	clone := original.Clone().(*prototype.ConcretePrototype)

	slog.Info("original and clone are different objects", slog.Bool("addr isSame", original == clone))
	slog.Info("data is clone", slog.Bool("data isSame", original.Data() == clone.Data()))
}
