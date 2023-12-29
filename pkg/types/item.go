package types

import (
	"fmt"
	"unicode"
)

type GoItem struct {
	Name    string
	Members map[MemberType]IMembers
}

func (str GoItem) IsPublic() bool {
	return unicode.IsUpper(rune(str.Name[0]))
}

func NewGoItem(name string) GoItem {
	str := GoItem{Name: name, Members: make(map[MemberType]IMembers)}
	return str
}

func (gi GoItem) Create(gm IMember) (GoItem, error) {
	members, exists := gi.Members[gm.GetType()]
	if !exists {
		members = NewGoMembers(gm.GetType())
	}

	updated, err := members.Add(gm)
	if err != nil {
		return gi, fmt.Errorf(" error: %s trying to add %s to %s", err.Error(), gm.GetName(), gi.Name)
	}
	gi.Members[gm.GetType()] = updated
	return gi, nil
}

func (gi GoItem) Read(mt MemberType, name string) (IMember, error) {
	members, exists := gi.Members[mt]
	if !exists {
		return nil, fmt.Errorf("%s does not have a member %s", mt.String(), name)
	}
	return members.GetMember(name)
}

func (gi GoItem) Update(gm IMember) (GoItem, error) {
	members, exists := gi.Members[gm.GetType()]
	if !exists {
		members = NewGoMembers(gm.GetType())
	}

	updated, err := members.Add(gm)
	if err != nil {
		return gi, fmt.Errorf(" error: %s trying to add %s to %s", err.Error(), gm.GetName(), gi.Name)
	}
	gi.Members[gm.GetType()] = updated
	return gi, nil
}

func (gi GoItem) Delete(mt MemberType, name string) (GoItem, error) {
	members, exists := gi.Members[mt]
	if !exists {
		return gi, nil
	}

	updated, err := members.Clear(name)
	if err != nil {
		return gi, fmt.Errorf(" error: %s trying to add %s to %s", err.Error(), name, gi.Name)
	}
	gi.Members[mt] = updated
	return gi, nil
}
