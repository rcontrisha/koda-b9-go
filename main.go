package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// perimeter, area := countRectangle(10, 5)
	// fmt.Printf("Perimeter: %d\n", perimeter)
	// fmt.Printf("Area: %d", area)
	// fmt.Println(countRectangle(10,5))

	// err := generateWindow(5)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	injectToSlice()
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

func generateWindow (width int) error {
	if width < 3 {
		return fmt.Errorf("Harus lebih dari 3")
	}

	for i := 1; i <= width; i++ {
		var window string = ""
		for j := 1; j <= width; j++ {
			if j == 1 || j == width || i == 1 || i == width {
				window += "*"
			} else {
				window += " "
			}
		}
		fmt.Println(window)
	}

	return nil
}

func injectToSlice () {
	var nums = []int8{50, 75, 66, 20, 32, 90}
	nums = append(nums[:4], nums[3:]...)
	fmt.Println(nums[:4])
	fmt.Println(nums[3:])
	fmt.Println(nums)
	nums[3] = 88

	fmt.Println(nums)
}