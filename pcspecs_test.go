package pcspecs

import (
	"fmt"
	"testing"
)

//
func TestSpec(t *testing.T) {
	expected := SysInfo{
		"DESKTOP-4DA5M71",
		"Microsoft Windows 10 Pro",
		"10",
		"Intel(R) Core(TM) i7-10700K CPU @ 3.80GHz",
		"NVIDIA GeForce RTX 2060 SUPER",
		"24",
		111,
		"TUF GAMING B460M-PLUS (WI-FI)",
	}
	got := Spec()
	fmt.Println(got)
	if got != expected {
		t.Error()
		t.Errorf("Expected: %v got: %v", expected, got)
	}
}
