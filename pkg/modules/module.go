package modules

import (
	"fmt"

	packages "github.com/fennysoftware/poe-core/pkg/packages"
)

type Module struct {
	Name string                      `json:"name"`
	pkgs map[string]packages.Package `json:"packages"`
}

func (m *Module) ListPackages() []string {
	names := []string{}
	for _, v := range m.pkgs {
		names = append(names, v.Name)
	}
	return names
}

func (m *Module) DelPackage(name string) error {
	_, ok := m.pkgs[name]
	if ok {
		return fmt.Errorf("cannot delete - package %s does not exists", name)
	} else {
		pkgs := m.pkgs
		delete(pkgs, name)
		m.pkgs = pkgs
	}
	return nil
}

func (m *Module) UpdatePackage(pkg packages.Package) {
	pkgs := m.pkgs
	pkgs[pkg.Name] = pkg
	m.pkgs = pkgs
}

func (m *Module) GetPackage(name string) (*packages.Package, error) {
	_, ok := m.pkgs[name]
	if !ok {
		return nil, fmt.Errorf("cannot get - package %s does not exists", name)
	}
	md := m.pkgs[name]
	return &md, nil
}

func (m *Module) CheckCreatingPackage(name string) error {
	_, ok := m.pkgs[name]
	if ok {
		return fmt.Errorf("cannot create - module %s already exists", name)
	} else {
		pkg := packages.NewPackage(name)
		m.pkgs[name] = pkg
	}
	return nil
}

func NewModule(name string) Module {
	return Module{Name: name, pkgs: make(map[string]packages.Package)}
}
