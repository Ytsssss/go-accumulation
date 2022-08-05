package algorithm

//func searchMatrix(matrix [][]int, target int) bool {
//	i, j := 0, 0
//	for i < len(matrix) {
//		if matrix[i][j] {
//
//		}
//	}
//}

func SearchMatrix74(matrix [][]int, target int) bool {
	i, height, length := 0, len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[height-1][length-1] {
		return false
	}
	k := height - 1
	middle := i + (k-i)/2
	for i <= k {
		middle = i + (k-i)/2
		if matrix[middle][0] == target {
			return true
		}
		if matrix[middle][0] < target {
			i = middle + 1
		} else if matrix[middle][0] > target {
			k = middle - 1
		}
	}

	i = k
	m, n := 0, len(matrix[i])-1
	for m <= n {
		middle1 := m + (n-m)/2
		if matrix[i][middle1] == target {
			return true
		}
		if matrix[i][middle1] < target {
			m = middle1 + 1
		} else if matrix[i][middle1] > target {
			n = middle1 - 1
		}
	}

	return false
}
