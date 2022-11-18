package pcspecs

import (
	"fmt"
	"testing"
)

//
func TestSpec(t *testing.T) {
	expected := SysInfo{
		"fv-az479-337",
		"Microsoft Windows Server 2022 Datacenter",
		"2022",
		"Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz",
		"Microsoft Hyper-V Video",
		"71",
		101,
		"Virtual Machine",
	}
	got := Spec()
	fmt.Println(got)
	if got != expected {
		t.Error()
		t.Errorf("Expected: %v got: %v", expected, got)
	}
}
