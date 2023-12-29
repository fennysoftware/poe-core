package packages

import (
	types "github.com/fennysoftware/poe-core/pkg/types"
)

const (
	MainPackageName = "main"
)

type Package struct {
	types.GoItem
}

func (pkg Package) IsMain(name string) bool {
	return IsMainPackage(pkg.Name)
}

func NewMainPackage() Package {
	return NewPackage(MainPackageName)
}

func NewPackage(name string) Package {
	return Package{GoItem: types.NewGoItem(name)}
}

func IsMainPackage(name string) bool {
	return name == MainPackageName
}

// Structures
func (pkg Package) StructureCreate(name string) (Package, error) {
	gi, err := pkg.GoItem.Create(types.NewGoMember(name, types.Mt_Structures))
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) StructureRead(name string) (types.IMember, error) {
	return pkg.GoItem.Read(types.Mt_Structures, name)
}

func (pkg Package) StructureUpdate(gm types.IMember) (Package, error) {
	gi, err := pkg.GoItem.Update(gm)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) StructureDelete(name string) (Package, error) {
	gi, err := pkg.GoItem.Delete(types.Mt_Structures, name)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

// Functions
func (pkg Package) FunctionCreate(name string) (Package, error) {
	gi, err := pkg.GoItem.Create(types.NewGoMember(name, types.Mt_Functions))
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) FunctionRead(name string) (types.IMember, error) {
	return pkg.GoItem.Read(types.Mt_Functions, name)
}

func (pkg Package) FunctionUpdate(gm types.IMember) (Package, error) {
	gi, err := pkg.GoItem.Update(gm)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) FunctionDelete(name string) (Package, error) {
	gi, err := pkg.GoItem.Delete(types.Mt_Functions, name)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

// Variables
func (pkg Package) VariableCreate(name string) (Package, error) {
	gi, err := pkg.GoItem.Create(types.NewGoMember(name, types.Mt_Variables))
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) VariableRead(name string) (types.IMember, error) {
	return pkg.GoItem.Read(types.Mt_Variables, name)
}

func (pkg Package) VariableUpdate(gm types.IMember) (Package, error) {
	gi, err := pkg.GoItem.Update(gm)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) VariableDelete(name string) (Package, error) {
	gi, err := pkg.GoItem.Delete(types.Mt_Variables, name)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

// Interfaces
func (pkg Package) InterfaceCreate(name string) (Package, error) {
	gi, err := pkg.GoItem.Create(types.NewGoMember(name, types.Mt_Interfaces))
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) InterfaceRead(name string) (types.IMember, error) {
	return pkg.GoItem.Read(types.Mt_Variables, name)
}

func (pkg Package) InterfaceUpdate(gm types.IMember) (Package, error) {
	gi, err := pkg.GoItem.Update(gm)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) InterfaceDelete(name string) (Package, error) {
	gi, err := pkg.GoItem.Delete(types.Mt_Interfaces, name)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

// Constants
func (pkg Package) ConstantCreate(name string) (Package, error) {
	gi, err := pkg.GoItem.Create(types.NewGoMember(name, types.Mt_Constants))
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) ConstantRead(name string) (types.IMember, error) {
	return pkg.GoItem.Read(types.Mt_Variables, name)
}

func (pkg Package) ConstantUpdate(gm types.IMember) (Package, error) {
	gi, err := pkg.GoItem.Update(gm)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}

func (pkg Package) ConstantDelete(name string) (Package, error) {
	gi, err := pkg.GoItem.Delete(types.Mt_Constants, name)
	if err != nil {
		return pkg, err
	}
	pkg.GoItem = gi
	return pkg, nil
}
