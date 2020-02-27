package pointer

import "testing"

func TestBasic(t *testing.T) {
	p := Person{Name: "玛丽"}

	p.SayName()
	p.SayName()
}
