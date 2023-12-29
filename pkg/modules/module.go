package modules

import (
	packages "github.com/fennysoftware/poe-core/pkg/packages"
)

type Module struct {
	Name     string
	Packages []packages.Package
}

func NewModule(name string) Module {
	return Module{Name: name, Packages: []packages.Package{}}
}
