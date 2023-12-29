package types

type MemberType int

const (
	Mt_Functions  MemberType = 0
	Mt_Variables  MemberType = 1
	Mt_Constants  MemberType = 2
	Mt_Structures MemberType = 3
	Mt_Interfaces MemberType = 4
)

const (
	// for errors and the like
	Mt_DisplayName_Function  = "Function"
	Mt_DisplayName_Variable  = "Variable"
	Mt_DisplayName_Constant  = "Constant"
	Mt_DisplayName_Structure = "Structure"
	Mt_DisplayName_Interface = "Interface"

	// for the enum
	Mt_DisplayName_Functions  = "Functions"
	Mt_DisplayName_Variables  = "Variables"
	Mt_DisplayName_Constants  = "Constants"
	Mt_DisplayName_Structures = "Structures"
	Mt_DisplayName_Interfaces = "Interfaces"
)

func (mt MemberType) String() string {
	switch mt {
	case Mt_Functions:
		return Mt_DisplayName_Functions
	case Mt_Variables:
		return Mt_DisplayName_Variables
	case Mt_Constants:
		return Mt_DisplayName_Constants
	case Mt_Structures:
		return Mt_DisplayName_Structures
	case Mt_Interfaces:
		return Mt_DisplayName_Interfaces
	}
	return "unknown"
}
