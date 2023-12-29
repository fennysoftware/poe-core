package packages

import (
	"testing"

	types "github.com/fennysoftware/poe-core/pkg/types"
)

const (
	strcname1  = "banana"
	strcnameu1 = "Banana"
)

func TestStructureAddVariable(t *testing.T) {
	str := NewStructure(strcname1)

	_, err := str.VariableCreate(strcname)
	if err != nil {
		t.Fatalf(`TestStructureAddVariable() Variable AddNew: %s error :%s`, str.Name, err.Error())
	}

	_, err = str.VariableRead(strcname)
	if err != nil {
		t.Fatalf(`TestStructureAddVariable() Member Get: %s error :%s`, strcname, err.Error())
	}
}

func TestStructureAddDuplicateVariable(t *testing.T) {
	str := NewStructure(strcname1)

	str, err := str.VariableCreate(strcname)
	if err != nil {
		t.Fatalf(`TestStructureAddDuplicateVariable() Variable AddNew: %s error :%s`, str.Name, err.Error())
	}

	str, err = str.VariableCreate(strcname)
	if err == nil {
		t.Fatal(`TestStructureAddDuplicateVariable() should have error but did not - must fail`)
	}

	mbrs := str.GoItem.Members

	pntrs := mbrs[types.Mt_Variables].GetAllPointers()
	if len(pntrs) > 0 {
		t.Fatal(`TestStructureAddDuplicateVariable() should have no pointers - must fail`)
	}

	refs := mbrs[types.Mt_Variables].GetAllReferences()
	if len(refs) == 0 {
		t.Fatal(`TestStructureAddDuplicateVariable() should have references - must fail`)
	}

}

/*
func TestMainPublicPrivateStruct(t *testing.T) {
	mn := NewMainPackage()
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`NewMainPackage() did not set package to main: %s`, mn.Name)
	}

	mn, err := AddNewStructure(mn, strcname)
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`AddNewStructure() failed with: %v`, err)
	}

	mn, err = AddNewStructure(mn, strcname)
	if err == nil {
		t.Fatal(`AddNewStructure() should have error but did not fail`)
	}

	str, exists := mn.Structures[strcname]
	if !exists {
		t.Fatal(`Structure Missing() failed`)
	}

	if str.IsPublic() {
		t.Fatal(`Structure should be private() failed`)
	}

	mn, err = AddNewStructure(mn, strcnameu)
	if err != nil {
		t.Fatal(`AddNewStructure() should have error but did not fail`)
	}

	mn, err = AddNewStructure(mn, strcnameu)
	if err == nil {
		t.Fatal(`AddNewStructure() should have error but did not fail`)
	}

	str, exists = mn.Structures[strcnameu]
	if !exists {
		t.Fatal(`Structure Missing() failed`)
	}

	if !str.IsPublic() {
		t.Fatal(`Structure should be private() failed`)
	}
}
*/
