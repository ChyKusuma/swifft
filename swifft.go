/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lswifft
#include "libswifft/swifft.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// AllocateInputBuffer allocates a buffer for SWIFFT input
func AllocateInputBuffer() []byte {
	size := C.SWIFFT_INPUT_BLOCK_SIZE
	buffer := C.malloc(C.size_t(size))
	if buffer == nil {
		panic("failed to allocate input buffer")
	}
	return C.GoBytes(buffer, C.int(size))
}

// AllocateOutputBuffer allocates a buffer for SWIFFT output
func AllocateOutputBuffer() []byte {
	size := C.SWIFFT_OUTPUT_BLOCK_SIZE
	buffer := C.malloc(C.size_t(size))
	if buffer == nil {
		panic("failed to allocate output buffer")
	}
	return C.GoBytes(buffer, C.int(size))
}

// AllocateCompactBuffer allocates a buffer for SWIFFT compact output
func AllocateCompactBuffer() []byte {
	size := C.SWIFFT_COMPACT_BLOCK_SIZE
	buffer := C.malloc(C.size_t(size))
	if buffer == nil {
		panic("failed to allocate compact buffer")
	}
	return C.GoBytes(buffer, C.int(size))
}

// ComputeHash computes the SWIFFT hash
func ComputeHash(input, output []byte) {
	C.SWIFFT_Compute(
		(*C.uchar)(unsafe.Pointer(&input[0])),
		(*C.uchar)(unsafe.Pointer(&output[0])),
	)
}

// CompactHash compacts the hash output
func CompactHash(output, compact []byte) {
	C.SWIFFT_Compact(
		(*C.uchar)(unsafe.Pointer(&output[0])),
		(*C.uchar)(unsafe.Pointer(&compact[0])),
	)
}
