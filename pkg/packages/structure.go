package packages

import types "github.com/fennysoftware/poe-core/pkg/types"

type Structure struct {
	types.GoItem
}

func NewStructure(name string) Structure {
	return Structure{GoItem: types.NewGoItem(name)}
}

// Variables
func (str Structure) VariableCreate(name string) (Structure, error) {
	gi, err := str.GoItem.Create(types.NewGoMember(name, types.Mt_Variables))
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

func (str Structure) VariableRead(name string) (types.IMember, error) {
	return str.GoItem.Read(types.Mt_Variables, name)
}

func (str Structure) VariableUpdate(gm types.IMember) (Structure, error) {
	gi, err := str.GoItem.Update(gm)
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

func (str Structure) VariableDelete(name string) (Structure, error) {
	gi, err := str.GoItem.Delete(types.Mt_Variables, name)
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

// Functions
func (str Structure) FunctionCreate(name string) (Structure, error) {
	gi, err := str.GoItem.Create(types.NewGoMember(name, types.Mt_Functions))
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

func (str Structure) FunctionRead(name string) (types.IMember, error) {
	return str.GoItem.Read(types.Mt_Functions, name)
}

func (str Structure) FunctionUpdate(gm types.IMember) (Structure, error) {
	gi, err := str.GoItem.Update(gm)
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

func (str Structure) FunctionDelete(name string) (Structure, error) {
	gi, err := str.GoItem.Delete(types.Mt_Functions, name)
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

// Structures
func (str Structure) StructureCreate(name string) (Structure, error) {
	gi, err := str.GoItem.Create(types.NewGoMember(name, types.Mt_Structures))
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

func (str Structure) StructureRead(name string) (types.IMember, error) {
	return str.GoItem.Read(types.Mt_Structures, name)
}

func (str Structure) StructureUpdate(gm types.GoMember) (Structure, error) {
	gi, err := str.GoItem.Update(gm)
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}

func (str Structure) StructureDelete(name string) (Structure, error) {
	gi, err := str.GoItem.Delete(types.Mt_Structures, name)
	if err != nil {
		return str, err
	}
	str.GoItem = gi
	return str, nil
}
