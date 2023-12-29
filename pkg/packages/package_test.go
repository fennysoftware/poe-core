package packages

import (
	"testing"
)

const (
	strcname  = "banana"
	strcnameu = "Banana"
)

func TestMain(t *testing.T) {
	mn := NewMainPackage()
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`NewMainPackage() did not set package to main: %s`, mn.Name)
	}
}

func TestMainAddStructure(t *testing.T) {
	mn := NewMainPackage()
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`NewMainPackage() did not set package to main: %s`, mn.Name)
	}

	mn, err := mn.StructureCreate(strcname)
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`StructureCreate() failed with: %v`, err)
	}

	_, err = mn.StructureRead(strcname)
	if err != nil {
		t.Fatal(`Structure Missing() failed`)
	}
}

func TestMainAddDuplicateStructure(t *testing.T) {
	mn := NewMainPackage()
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`NewMainPackage() did not set package to main: %s`, mn.Name)
	}

	mn, err := mn.StructureCreate(strcname)
	if err != nil {
		t.Fatalf(`CreateStructure() failed with: %v`, err)
	}

	mn, err = mn.StructureCreate(strcname)
	if err == nil {
		t.Fatal(`CreateStructure() should have error but did not fail`)
	}
}

func TestMainPublicPrivateStruct(t *testing.T) {
	mn := NewMainPackage()
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`NewMainPackage() did not set package to main: %s`, mn.Name)
	}

	mn, err := mn.StructureCreate(strcname)
	if !IsMainPackage(mn.Name) {
		t.Fatalf(`CreateStructure() failed with: %v`, err)
	}

	mn, err = mn.StructureCreate(strcname)
	if err == nil {
		t.Fatal(`CreateStructure() should have error but did not fail`)
	}

	str, err := mn.StructureRead(strcname)
	if err != nil {
		t.Fatal(`Structure Missing() failed`)
	}

	if str.IsPublic() {
		t.Fatal(`Structure should be private() failed`)
	}

	mn, err = mn.StructureCreate(strcnameu)
	if err != nil {
		t.Fatal(`CreateStructure() should have error but did not fail`)
	}

	mn, err = mn.StructureCreate(strcnameu)
	if err == nil {
		t.Fatal(`CreateStructure() should have error but did not fail`)
	}

	str, err = mn.StructureRead(strcnameu)
	if err != nil {
		t.Fatal(`Structure Missing() failed`)
	}

	if !str.IsPublic() {
		t.Fatal(`Structure should be private() failed`)
	}
}
