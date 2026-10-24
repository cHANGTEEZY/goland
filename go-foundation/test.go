package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// signed integers
	var i8 int8 = 127             // range: -128 to 127
	var i16 int16 = 32_767        // range: -32,768 to 32,767
	var i32 int32 = 2_147_483_647 // range: -2^31 to 2^31-1
	var i64 int64 = math.MaxInt64 // range: -2^63 to 2^63-1

	fmt.Println("signed Integers")
	fmt.Printf("int8   : %d  (size: %d byte)\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("int16  : %d  (size: %d bytes)\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("int32  : %d  (size: %d bytes)\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("int64  : %d  (size: %d bytes)\n", i64, unsafe.Sizeof(i64))

	// unsigned integers
	var u8 uint8 = 255              // range: 0 to 255
	var u16 uint16 = 65_535         // range: 0 to 65,535
	var u32 uint32 = 4_294_967_295  // range: 0 to 2^32-1
	var u64 uint64 = math.MaxUint64 // range: 0 to 2^64-1

	fmt.Println("\n unsigned Integers ")
	fmt.Printf("uint8  : %d  (size: %d byte)\n", u8, unsafe.Sizeof(u8))
	fmt.Printf("uint16 : %d  (size: %d bytes)\n", u16, unsafe.Sizeof(u16))
	fmt.Printf("uint32 : %d  (size: %d bytes)\n", u32, unsafe.Sizeof(u32))
	fmt.Printf("uint64 : %d  (size: %d bytes)\n", u64, unsafe.Sizeof(u64))

	var i int = math.MaxInt
	var u uint = math.MaxUint

	fmt.Println("\n platform-Dependent ")
	fmt.Printf("int    : %d  (size: %d bytes)\n", i, unsafe.Sizeof(i))
	fmt.Printf("uint   : %d  (size: %d bytes)\n", u, unsafe.Sizeof(u))

	var b byte = 'A'
	fmt.Println("\n special synonyms ")
	fmt.Printf("byte   : %d  '%c'  (alias for uint8,  size: %d byte)\n", b, b, unsafe.Sizeof(b))

	var r rune = '🎉'
	fmt.Printf("rune   : %d  '%c'  (alias for int32, size: %d bytes)\n", r, r, unsafe.Sizeof(r))

	fmt.Println()

	// uintptr
	x := 42
	ptr := uintptr(unsafe.Pointer(&x)) // convert pointer to uintptr for arithmetic

	fmt.Println("\n low-Level: uintptr ")
	fmt.Printf("uintptr: 0x%X  (size: %d bytes, holds pointer to x=%d)\n",
		ptr, unsafe.Sizeof(ptr), x)

	// overflow demonstration
	var overflow uint8 = 255
	overflow++ // wraps around to 0

	var underflow int8 = -128
	underflow-- // wraps around to 127

	fmt.Println("\noverflow / underflow (wrap-around)")
	fmt.Printf("uint8(255)++ = %d  (wrapped to 0)\n", overflow)
	fmt.Printf("int8(-128)-- = %d  (wrapped to 127)\n", underflow)
}
