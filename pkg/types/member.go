package types

import (
	"unicode"
)

type GoMember struct {
	IMember
	name          string
	pointerPrefix bool
	mtype         MemberType
}

func (gm GoMember) GetName() string {
	return gm.name
}

func (gm GoMember) SetName(name string) IMember {
	gm.name = name
	return gm
}

func (gm GoMember) GetType() MemberType {
	return gm.mtype
}

func (gm GoMember) SetType(mtype MemberType) IMember {
	gm.mtype = mtype
	return gm
}

func (gm GoMember) HasPointerPrefix() bool {
	return gm.pointerPrefix
}

func (gm GoMember) SetPointerPrefix(ptr bool) IMember {
	gm.pointerPrefix = ptr
	return gm
}

func (gm GoMember) IsPublic() bool {
	return unicode.IsUpper(rune(gm.name[0]))
}

func NewGoMember(name string, t MemberType) GoMember {
	gm := GoMember{name: name, mtype: t, pointerPrefix: false}
	return gm
}

func NewGoPtrMember(name string, t MemberType) GoMember {
	gm := GoMember{name: name, mtype: t, pointerPrefix: true}
	return gm
}
