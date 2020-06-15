package structfieldorder

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMyStruct(t *testing.T) {
	a := MyStruct{}
	aSize := unsafe.Sizeof(a) // 24 bytes

	b := MyStructOptimized{}
	bSize := unsafe.Sizeof(b) // ordered 16 bytes

	if bSize >= aSize {
		t.Errorf("optimized struct should be less than unoptimized")
	}

	fmt.Printf("unoptimized size: %v bytes  optimized size: %v bytes\n", aSize, bSize)
}
