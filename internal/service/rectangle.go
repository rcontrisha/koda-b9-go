package service

func CountPerimeter (width uint8, height uint8) uint8 {
	return 2 * (width + height)
}

func CountArea (width uint8, height uint8) uint8 {
	return width * height
}

func CountRectangle (width uint8, height uint8) (perimeter uint8, area uint8)  {
	perimeter = CountPerimeter(width, height)	
	area = CountArea(width, height)

	return perimeter, area
}
