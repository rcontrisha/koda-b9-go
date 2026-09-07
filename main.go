package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	perimeter, area := countRectangle(10, 5)
	fmt.Printf("Perimeter: %d\n", perimeter)
	fmt.Printf("Area: %d", area)
	// fmt.Println(countRectangle(10,5))
}

func countPerimeter (width uint8, height uint8) uint8 {
	return 2 * (width + height)
}

func countArea (width uint8, height uint8) uint8 {
	return width * height
}

func countRectangle (width uint8, height uint8) (perimeter uint8, area uint8)  {
	perimeter = countPerimeter(width, height)
	area = countArea(width, height)

	return perimeter, area
}