package piece

import "testing"

func TestEnum(t *testing.T) {
	if King != 1 {
		t.Fatalf("King should == 1")
	}
}
