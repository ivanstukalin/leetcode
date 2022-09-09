package main

//O(n)
func spiralOrder(matrix [][]int) []int {
	var result []int
	var countElements = len(matrix[0]) * len(matrix)
	currentCount := 0
	position := 1

	for currentCount < countElements {
		columnLength := len(matrix[0]) - position
		rowsLength := len(matrix) - position
		for key, value := range matrix[position-1] {
			if key >= position-1 && key <= columnLength {
				result = append(result, value)
				currentCount++
			}
		}
		for key, value := range matrix {
			if key > position-1 && key <= rowsLength && currentCount < countElements {
				result = append(result, value[columnLength])
				currentCount++
			}
		}
		if len(matrix) > 1 {
			for i := columnLength; i >= 0; i-- {
				if i != columnLength && i >= position-1 && currentCount < countElements {
					result = append(result, matrix[rowsLength][i])
					currentCount++
				}
			}
		}
		if len(matrix) != 2 && len(matrix) != 1 {
			for i := rowsLength; i >= position; i-- {
				if i != rowsLength && i != position-1 && currentCount < countElements {
					result = append(result, matrix[i][position-1])
					currentCount++
				}
			}
		}
		position++
	}
	return result
}
