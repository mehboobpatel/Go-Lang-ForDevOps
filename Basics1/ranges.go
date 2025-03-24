package main

import "fmt"

func main() {
	// Signed Integers
	// int8: 8-bit signed integer (-128 to 127)
	var a int8 = 127 // Last value in range

	// int16: 16-bit signed integer (-32768 to 32767)
	var b int16 = 32767 // Last value in range

	// int32: 32-bit signed integer (-2147483648 to 2147483647)
	var c int32 = 2147483647 // Last value in range

	// int64: 64-bit signed integer (-9223372036854775808 to 9223372036854775807)
	var d int64 = 9223372036854775807 // Last value in range

	// Unsigned Integers
	// uint8: 8-bit unsigned integer (0 to 255)
	var e uint8 = 255 // Last value in range

	// uint16: 16-bit unsigned integer (0 to 65535)
	var f uint16 = 65535 // Last value in range

	// uint32: 32-bit unsigned integer (0 to 4294967295)
	var g uint32 = 4294967295 // Last value in range

	// uint64: 64-bit unsigned integer (0 to 18446744073709551615)
	var h uint64 = 18446744073709551615 // Last value in range

	// Floating-Point Numbers
	// float32: 32-bit floating-point number (~6 decimal digits of precision)
	var i float32 = 3.40282346638528859811704183484516925440e+38 // Largest value for float32

	// float64: 64-bit floating-point number (~15 decimal digits of precision)
	var j float64 = 1.797693134862315708145274237317043567981e+308 // Largest value for float64

	// Complex Numbers
	// complex64: 32-bit real and imaginary parts
	var k complex64 = complex(3.40282346638528859811704183484516925440e+38, 3.40282346638528859811704183484516925440e+38) // Largest value for complex64

	// complex128: 64-bit real and imaginary parts
	var l complex128 = complex(1.797693134862315708145274237317043567981e+308, 1.797693134862315708145274237317043567981e+308) // Largest value for complex128

	// Print the values
	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)
	fmt.Println("uint8:", e)
	fmt.Println("uint16:", f)
	fmt.Println("uint32:", g)
	fmt.Println("uint64:", h)
	fmt.Println("float32:", i)
	fmt.Println("float64:", j)
	fmt.Println("complex64:", k)
	fmt.Println("complex128:", l)
}
