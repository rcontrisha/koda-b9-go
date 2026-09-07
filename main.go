package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	countRectangle(10, 5)
}

func countPerimeter (width uint8, height uint8) uint8 {
	return 2 * (width + height)
}

func countArea (width uint8, height uint8) uint8 {
	return width * height
}

func countRectangle (width uint8, height uint8) {
	var perimeter uint8 = countPerimeter(width, height)
	var area uint8 = countArea(width, height)

	fmt.Println("\nCount Area & Perimeter of Rectangle")
	fmt.Println("=======================================")
	fmt.Printf("Area: %d\n", area)
	fmt.Printf("Perimeter: %d", perimeter)
}