package types

type IMember interface {
	IsPublic() bool
	HasPointerPrefix() bool
	SetPointerPrefix(bool) IMember
	GetName() string
	SetName(string) IMember
	GetType() MemberType
	SetType(MemberType) IMember
}
