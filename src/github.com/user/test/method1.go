package main

import (
	"fmt"
)

// Retangle xx
type Retangle struct {
	width, height float64
}

// GetWidth return Retangle width
func (area Retangle) GetWidth() (width float64) {
	width = area.width
	return
}

// GetHeight return Retangle height
func (area *Retangle) GetHeight() float64 {

	return (*area).height
}

func main() {
	var r1 = Retangle{12, 13}
	fmt.Println("Area With of r1 is : ", r1.GetWidth())
	fmt.Println("Area Height of r1 is : ", r1.GetHeight())
}
