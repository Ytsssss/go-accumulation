package algorithm

func SolveNQueens(n int) [][]string {
	return dfs(0, n, []string{}, [][]string{}, map[int]bool{}, map[int]bool{}, map[int]bool{})
}

func dfs(depth, n int, current []string, result [][]string, lie, pie, la map[int]bool) [][]string {
	if n == depth {
		return append(result, current)
	}
	for j := 0; j < n; j++ {
		str := ""
		if lie[j] || pie[j+depth] || la[depth-j] {
			continue
		}
		for i := 0; i < n; i++ {
			if i == j {
				str += "Q"
			} else {
				str += "."
			}
		}
		tmp := make([]string, 0)
		tmp = append(tmp, current...)
		lie[j], pie[j+depth], la[depth-j] = true, true, true
		result = dfs(depth+1, n, append(current, str), result, lie, pie, la)
		current = tmp
		delete(lie, j)
		delete(pie, j+depth)
		delete(la, depth-j)
	}
	return result
}

func TotalNQueens(n int) int {
	// col,pie,la 分别表示列/左斜对角/右斜对角被占用情况
	col, pie, la := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	return getCount(n, col, pie, la, 0, 0)
}

func getCount(n int, col, pie, la map[int]bool, depth, count int) int {
	if depth == n {
		return count + 1
	}
	for i := 0; i < n; i++ {
		if col[i] || pie[i+depth] || la[i-depth] {
			continue
		}
		col[i], pie[i+depth], la[i-depth] = true, true, true
		count = getCount(n, col, pie, la, depth+1, count)
		col[i], pie[i+depth], la[i-depth] = false, false, false
	}
	return count
}

func TotalNQueens2(n int) int {
	// col,pie,la 分别表示列/左斜对角/右斜对角被占用情况
	col, pie, la := 0, 0, 0
	return getCount2(n, col, pie, la, 0, 0)
}

func getCount2(n, col, pie, la, depth, count int) int {
	if depth == n {
		return count + 1
	}
	for i := 0; i < n; i++ {
		if i&(col|pie|la) > 0 {
			continue
		}
		tmpCol, tmpPie, tmpLa := col, pie, la
		col, pie, la = col|i, (pie|i)<<1, (la|i)>>1
		count = getCount2(n, col, pie, la, depth+1, count)
		col, pie, la = tmpCol, tmpPie, tmpLa
	}
	return count
}
