package algorithm

func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	answer := make([]int, 0, m*n)
	i, j, count := 0, 0, 0
	for i < m && j < n && count <= m*n {
		for k := j; k < n-j; k++ {
			answer = append(answer, matrix[i][k])
			count++
		}
		for k := i + 1; k < m-i; k++ {
			answer = append(answer, matrix[k][n-j-1])
			count++
		}
		for k := n - j - 2; k >= j && count <= m*n; k-- {
			answer = append(answer, matrix[m-i-1][k])
			count++
		}
		for k := m - i - 2; k >= i+1 && count <= m*n; k-- {
			answer = append(answer, matrix[k][i])
			count++
		}
		i++
		j++
	}
	return answer
}
