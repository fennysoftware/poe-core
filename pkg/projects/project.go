package projects

import (
	"fmt"

	modules "github.com/fennysoftware/poe-core/pkg/modules"
)

type Project struct {
	Name string                    `json:"name"`
	mods map[string]modules.Module `json:"modules"`
}

func (p *Project) ListModules() []string {
	names := []string{}
	for _, v := range p.mods {
		names = append(names, v.Name)
	}
	return names
}

func (p *Project) DelModule(modname string) error {
	_, ok := p.mods[modname]
	if ok {
		return fmt.Errorf("cannot delete - module %s does not exists", modname)
	} else {
		mods := p.mods
		delete(mods, modname)
		p.mods = mods
	}

	return nil
}

func (p *Project) UpdateModule(mod modules.Module) {
	mods := p.mods
	mods[mod.Name] = mod
	p.mods = mods
}

func (p *Project) GetModule(modname string) (*modules.Module, error) {
	_, ok := p.mods[modname]
	if !ok {
		return nil, fmt.Errorf("cannot get - module %s does not exists", modname)
	}
	md := p.mods[modname]
	return &md, nil
}

func (p *Project) CheckCreatingModule(modname string) error {
	_, ok := p.mods[modname]
	if ok {
		return fmt.Errorf("cannot create - module %s already exists", modname)
	} else {
		md := modules.NewModule(modname)
		p.mods[modname] = md
	}
	return nil
}

func NewProject(name string) *Project {
	return &Project{Name: name, mods: make(map[string]modules.Module)}
}
