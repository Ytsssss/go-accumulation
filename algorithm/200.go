package algorithm

func NumIslands(grid [][]byte) int {
	island := &Island{
		Grid: grid,
	}
	island.initIsland()
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i < len(grid)-1 && grid[i+1][j] == '1' {
					island.union(i*n+j, (i+1)*n+j)
				}
				if j < len(grid[i])-1 && grid[i][j+1] == '1' {
					island.union(i*n+j, i*n+j+1)
				}
				if i > 0 && grid[i-1][j] == '1' {
					island.union(i*n+j, (i-1)*n+j)
				}
				if j > 0 && grid[i][j-1] == '1' {
					island.union(i*n+j, i*n+j-1)
				}
			}
		}
	}
	return island.Answer
}

type Island struct {
	Nums   []int
	Grid   [][]byte
	Answer int
	Rank   []int
}

func (i *Island) findRoot(k int) int {
	if i.Nums[k] != k {
		i.Nums[k] = i.findRoot(i.Nums[k])
	}
	return i.Nums[k]
}

func (i *Island) initIsland() {
	m, n := len(i.Grid), len(i.Grid[0])
	nums := make([]int, m*n)
	rank := make([]int, m*n)
	i.Nums, i.Rank = nums, rank
	for k := 0; k < m; k++ {
		for j := 0; j < n; j++ {
			if i.Grid[k][j] == '1' {
				i.Answer++
				i.Nums[j+n*k] = j + n*k
			}
			i.Rank[j+k*n] = 0
		}
	}
}

func (i *Island) union(k, j int) {
	kroot := i.findRoot(k)
	jroot := i.findRoot(j)
	if kroot != jroot {
		if i.Rank[kroot] > i.Rank[jroot] {
			i.Nums[jroot] = kroot
		} else if i.Rank[kroot] < i.Rank[jroot] {
			i.Nums[kroot] = jroot
		} else {
			i.Nums[jroot] = kroot
			i.Rank[kroot]++
		}
		i.Answer--
	}
}
