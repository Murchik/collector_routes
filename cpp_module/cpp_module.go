package cpp_module

/*
#cgo CFLAGS: -I./cpp_module
#cgo LDFLAGS: -L./cpp_module -lAlgorithms

#include "Algorithms.hpp"
*/
import "C"

func SayHello() {
	// Test int hello();
	C.hello()
}
