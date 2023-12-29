package types

import (
	"fmt"
)

type GoMembers struct {
	IMembers
	mtype   MemberType
	members map[string]IMember
}

func NewGoMembers(mt MemberType) IMembers {
	return GoMembers{mtype: mt, members: make(map[string]IMember)}
}

func (gms GoMembers) GetType() MemberType {
	return gms.mtype
}

func (gms GoMembers) SetType(mtype MemberType) IMembers {
	gms.mtype = mtype
	return gms
}

func (gms GoMembers) Add(memb IMember) (IMembers, error) {
	if memb.GetType() != gms.mtype {
		return gms, fmt.Errorf("%s does not match %s - for %s", gms.mtype.String(), memb.GetType().String(), memb.GetName())
	}
	members := gms.members
	_, exists := members[memb.GetName()]
	if exists {
		return gms, fmt.Errorf("%s %s already exists", gms.mtype.String(), memb.GetName())
	}
	members[memb.GetName()] = memb
	gms.members = members
	return gms, nil
}

func (gms GoMembers) Update(memb IMember) (IMembers, error) {
	if memb.GetType() != gms.mtype {
		return gms, fmt.Errorf("%s does not match %s - for %s", gms.mtype.String(), memb.GetType().String(), memb.GetName())
	}
	members := gms.members
	_, exists := members[memb.GetName()]
	if !exists {
		return gms.Add(memb)
	}
	members[memb.GetName()] = memb
	gms.members = members
	return gms, nil
}

func (gms GoMembers) Clear(name string) (IMembers, error) {
	members := gms.members
	_, exists := members[name]
	if !exists {
		return gms, nil
	}
	delete(members, name)
	gms.members = members
	return gms, nil
}

func (gms GoMembers) GetMember(name string) (IMember, error) {
	member, exists := gms.members[name]
	if !exists {
		return nil, fmt.Errorf("member %s missing", name)
	}
	return member, nil
}

func (gms GoMembers) GetAllPointers() map[string]IMember {

	results := make(map[string]IMember)
	for name, v := range gms.members {
		if v.HasPointerPrefix() {
			results[name] = v
		}
	}
	return results
}

func (gms GoMembers) GetAllReferences() map[string]IMember {
	results := make(map[string]IMember)
	for name, v := range gms.members {
		if !v.HasPointerPrefix() {
			results[name] = v
		}
	}
	return results
}
