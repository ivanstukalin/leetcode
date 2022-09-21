package main

func setZeroes(matrix [][]int) [][]int  {
	col := 0
	for rowId, row := range matrix {
		if row[0] == 0 {
			col = 1
		}
		for i := 1; i < len(row); i++ {
			if row[i] == 0 {
				matrix[0][i] = 0
				matrix[rowId][0] = 0
			}
		}
	}
	for rowId := len(matrix) - 1; rowId >= 0; rowId-- {
		for columnId := len(matrix[rowId]) - 1; columnId >= 1; columnId-- {
			if matrix[rowId][0] == 0 || matrix[0][columnId] == 0 {
				matrix[rowId][columnId] = 0
			}
		}
		if col == 1 {
			matrix[rowId][0] = 0
		}
	}
	return matrix
}
