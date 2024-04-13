// Example of syscolor package usage
package main

import (
	"fmt"

	"github.com/mishamyrt/go-syscolor"
)

func main() {
	fg, err := syscolor.SelectedForeground()
	if err != nil {
		panic(err)
	}
	bg, err := syscolor.SelectedBackground()
	if err != nil {
		panic(err)
	}
	fmt.Printf("foreground: %v\nbackground: %v\n", fg, bg)
}
