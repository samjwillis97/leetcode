package floodFill

import "fmt"

// An image is represented by an m x n integer grid image
// where image[i][j] represents the pixel value of the image.
//
// You are also given three integers sr, sc, and newColor.
// You should perform a flood fill on the image starting from the pixel image[sr][sc].
//
// To perform a flood fill, consider the starting pixel, plus any pixels connected
// 4-directionally to the starting pixel of the same color as the starting pixel,
// plus any pixels connected 4-directionally to those pixels (also with the same color),
// and so on. Replace the color of all the aforementioned pixels with newColor.
//
// Return the modified image after performing the flood fill.

// Recursive
// Current, then Neighbouring

// SR = 1, SC = 1, NC = 2
// Image
// 1 1 1	  2 2 2
// 1 1 0  ->  2 2 0
// 1 0 1	  2 0 1
// From the center of the image with position (sr, sc) = (1, 1)
// (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel
// (i.e., the blue pixels) are colored with the new color.
// Note the bottom corner is not colored 2,
// because it is not 4-directionally connected to the starting pixel.

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// Start at SR, SC
	// i = rows
	// j = cols
	iBound := len(image)
	if iBound == 0 {
		return image
	}
	if image[sr][sc] == newColor {
		return image
	}

	image = colorPixels(image, sr, sc, newColor)

	return image
}

func colorPixels(image [][]int, i int, j int, newColor int) [][]int {
	mainColor := image[i][j]
	//fmt.Printf("i: %d, j: %d\n", i, j)
	image[i][j] = newColor
	//printImage(image)
	iBound := len(image) - 1
	jBound := len(image[0]) - 1
	if i < iBound {
		iNew := i + 1
		if image[iNew][j] == mainColor {
			//fmt.Println("V")
			image = colorPixels(image, iNew, j, newColor)
		}
	}
	if i > 0 {
		iNew := i - 1
		if image[iNew][j] == mainColor {
			//fmt.Println("^")
			image = colorPixels(image, iNew, j, newColor)
		}
	}
	if j < jBound {
		jNew := j + 1
		if image[i][jNew] == mainColor {
			//fmt.Println("->")
			image = colorPixels(image, i, jNew, newColor)
		}
	}
	if j > 0 {
		jNew := j - 1
		if image[i][jNew] == mainColor {
			//fmt.Println("<-")
			image = colorPixels(image, i, jNew, newColor)
		}
	}

	return image
}

func printImage(a [][]int) {
	for _, row := range a {
		for _, i := range row {
			fmt.Printf("|%d", i)
		}
		fmt.Printf("|\n")
	}
}
