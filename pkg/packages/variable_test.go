package packages

import (
	"testing"
)

const (
	vartest = "somevar"
)

func TestVariableAddVariable(t *testing.T) {
	str := NewStructure(strcname1)

	v := NewRefVariable(vartest, "int")
	_, err := str.VariableUpdate(v)
	if err != nil {
		t.Fatalf(`TestStructureAddVariable() Variable AddNew: %s error :%s`, str.Name, err.Error())
	}

	vi, err := str.VariableRead(vartest)
	if err != nil {
		t.Fatalf(`TestStructureAddVariable() Member Get: %s error :%s`, strcname, err.Error())
	}

	updated := vi.(Variable)

	if updated.GetName() != vartest {
		t.Fatalf(`Variable is called: %s of type :%s is public %t`, updated.GetName(), updated.Type, updated.IsPublic())
	}

	if updated.Type != "int" {
		t.Fatalf(`Variable is called: %s of type :%s is public %t`, updated.GetName(), updated.Type, updated.IsPublic())
	}
}
