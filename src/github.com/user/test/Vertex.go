package test

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Say the args input
func Say(s string) {
	fmt.Printf("s=%s\n", s)
}

// AppendByte append byte autolly
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > m {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
