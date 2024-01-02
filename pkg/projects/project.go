package projects

import (
	"fmt"

	modules "github.com/fennysoftware/poe-core/pkg/modules"
)

type Project struct {
	Name string
	mods map[string]modules.Module
}

func (p *Project) CheckCreatingModule(modname string) {
	_, ok := p.mods[modname]
	if ok {
		fmt.Printf("Module %s already exists.\n", modname)
	} else {
		md := modules.NewModule(modname)
		p.mods[modname] = md
		fmt.Printf("Module %s created.\n", modname)
	}
}

func NewProject(name string) *Project {
	return &Project{Name: name, mods: make(map[string]modules.Module)}
}
