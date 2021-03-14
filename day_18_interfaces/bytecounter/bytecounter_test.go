package bytecounter

import (
	"fmt"
	"testing"
)

func TestByCounterWrite(t *testing.T) {
	name := "Quyen"
	var c ByteCounter
	if got, _ := fmt.Fprintf(&c, "%s", name); got != len(name) {
		t.Errorf(`len of %s = %d, want %d`, name, got, len(name))
	}
}
