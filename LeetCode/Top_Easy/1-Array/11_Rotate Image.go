package main

func main() {

	// input := [4][4]int{
	// 	{0, 1, 2, 3},
	// 	{4, 5, 6, 7},
	// 	{8, 9, 10, 11},
	// 	{12, 13, 14, 15}}

	// rotate(input)
	// for i:=0; i<len(input);i++{
	// 	fmt.Println(input[i])
	// }

}

// Group Cell Apporach
func rotate(matrix [][]int) {
	sizeA := len(matrix)
	for i := 0; i < sizeA/2; i++ {
		for j := i; j < sizeA-(i+1); j++ {

			// 1- matrix[i][j]
			// 2- matrix[j][(sizeA-1)-i]
			// 3- matrix[(sizeA-1)-i][(sizeA-1)-j]
			// 4- matrix[(sizeA-1)-j][(sizeA-1)-((sizeA-1)-i)] -> matrix[(sizeA-1)-j][i]

			matrix[i][j], matrix[(sizeA-1)-j][i] = matrix[(sizeA-1)-j][i], matrix[i][j]                                         //1 <-> 4
			matrix[(sizeA-1)-j][i], matrix[(sizeA-1)-i][(sizeA-1)-j] = matrix[(sizeA-1)-i][(sizeA-1)-j], matrix[(sizeA-1)-j][i] //4 <-> 3
			matrix[(sizeA-1)-i][(sizeA-1)-j], matrix[j][(sizeA-1)-i] = matrix[j][(sizeA-1)-i], matrix[(sizeA-1)-i][(sizeA-1)-j] //3 <-> 2

		}
	}
}

// Tanspose Reverse Approach
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

		for x, y := 0, len(matrix)-1; x < y; x, y = x+1, y-1 {
			matrix[i][x], matrix[i][y] = matrix[i][y], matrix[i][x]
		}
	}
}
