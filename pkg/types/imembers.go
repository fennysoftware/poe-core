package types

type IMembers interface {
	GetType() MemberType
	SetType(MemberType) IMembers
	Add(IMember) (IMembers, error)
	Update(IMember) (IMembers, error)
	Clear(string) (IMembers, error)
	GetMember(string) (IMember, error)
	GetAllPointers() map[string]IMember
	GetAllReferences() map[string]IMember
}
