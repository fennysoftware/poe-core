package packages

import types "github.com/fennysoftware/poe-core/pkg/types"

type Variable struct {
	types.GoMember
	Type string
}

func NewRefVariable(name string, t string) types.IMember {
	return Variable{GoMember: types.NewGoMember(name, types.Mt_Variables), Type: t}
}

func NewPtrVariable(name string, t string) types.IMember {
	return Variable{GoMember: types.NewGoPtrMember(name, types.Mt_Variables), Type: t}
}
