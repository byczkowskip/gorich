package main

import (
	"fmt"

	"github.com/byczkowskip/gorich.git/rich"
)

func main() {
	// using styles
	rich.Println("Bold + Underline", rich.Bold, rich.Underline)

	// using a color by name
	red := rich.FromColorName("red")
	rich.Println("This is red text", red)

	// using colors and styles
	green := rich.FromHex("#00ff00")
	rich.Println("Green and bold", green, rich.Bold)

	// using the Printf function
	rich.Printf("Hello, %s!\n", []any{"world"}, red, rich.Underline)

	// using the Sprint function
	msg := rich.Sprint("Returned string", red, rich.Bold)
	fmt.Println("Normal fmt.Println, but string is styled ->", msg)

	// using the Sprintf function
	formatted := rich.Sprintf("Pi ≈ %.2f", []any{3.14159}, green)
	fmt.Println(formatted)
}
