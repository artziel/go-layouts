package layouts

import (
	"testing"
)

func TestLayoutStruct(t *testing.T) {

	l := NewLayout(TestRow{})
	l.AddRow(TestRow{})
	l.AddRow(TestRow{})
	l.AddRow(TestRow{})
	l.AddRow(TestRow{})
	l.AddRow(TestRow{})
	l.AddError(Error{Error: ErrRequiredValueRuleFail, Column: "A"})
	l.AddError(Error{Error: ErrRequiredValueRuleFail, Column: "B"})
	l.AddError(Error{Error: ErrRequiredValueRuleFail, Column: "C"})
	l.AddError(Error{Error: ErrRequiredValueRuleFail, Column: "D"})
	l.AddError(Error{Error: ErrRequiredValueRuleFail, Column: "E"})

	errs := l.GetErrors()
	if len(errs) != 5 {
		t.Errorf("Test 0: GetErrors should return 5 errors, Return %d \n", len(errs))
	}

	rows := l.GetRows()
	if len(rows) != 5 {
		t.Errorf("Test 1: GetRows should return 5 rows, Return %d\n", len(rows))
	}

}
