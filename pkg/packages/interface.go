package packages

import types "github.com/fennysoftware/poe-core/pkg/types"

type Interface struct {
	types.GoItem
}

func NewInterface(name string) Interface {
	return Interface{GoItem: types.NewGoItem(name)}
}

// Functions
func (intr Interface) FunctionCreate(name string) (Interface, error) {
	gi, err := intr.GoItem.Create(types.NewGoMember(name, types.Mt_Functions))
	if err != nil {
		return intr, err
	}
	intr.GoItem = gi
	return intr, nil
}

func (intr Interface) FunctionRead(name string) (types.IMember, error) {
	return intr.GoItem.Read(types.Mt_Functions, name)
}

func (intr Interface) FunctionUpdate(gm types.IMember) (Interface, error) {
	gi, err := intr.GoItem.Update(gm)
	if err != nil {
		return intr, err
	}
	intr.GoItem = gi
	return intr, nil
}

func (intr Interface) FunctionDelete(name string) (Interface, error) {
	gi, err := intr.GoItem.Delete(types.Mt_Functions, name)
	if err != nil {
		return intr, err
	}
	intr.GoItem = gi
	return intr, nil
}
