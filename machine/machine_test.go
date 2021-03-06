package machine

import (
	"testing"
)

func TestNew(t *testing.T) {
	m1 := New("XXX")
	m2 := Machine{"XXX"}

	if *m1 != m2 {
		t.Error("machine.New factory failed to produce appropriate machine.Machine")
	}

	if m1.BootId != "XXX" {
		t.Fatal("machine.Machine.BootId != 'XXX'")
	}
}

func TestStringEncoding(t *testing.T) {
	m1 := Machine{"XXX"}
	result := m1.String()
	expect := "XXX"
	if result != expect {
		t.Fatalf("machine.Machine.String() returned '%s', expected '%s'", result, expect)
	}
}
