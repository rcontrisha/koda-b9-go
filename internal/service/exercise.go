package service

import "fmt"

func GenerateWindow (width int) error {
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

func InjectToSlice () {
	var nums = []int8{50, 75, 66, 20, 32, 90}
	
	fmt.Println("Slice Before Injecting Element")
	fmt.Println("=============================")

	for i := range nums {
		fmt.Println(nums[i])
	}

	nums = append(nums[:4], nums[3:]...)
	// fmt.Println(nums[:4])
	// fmt.Println(nums[3:])
	// fmt.Println(nums)
	nums[3] = 88

	fmt.Println("\nSlice After Injecting Element")
	fmt.Println("=============================")

	for i := range nums {
		fmt.Println(nums[i])
	}
}