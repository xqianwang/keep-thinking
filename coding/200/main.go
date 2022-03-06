package main

import "fmt"

func main()  {
	//test := [][]byte{{'1','1','1','1','0'}, {'1','1','0','1','0'}, {'1','1','0','0','0'}, {'0','0','0','0','0'}}
	test := [][]byte{{'1','1'},{'1','1'}}
	fmt.Println(numIslands(test))
}

func numIslands(grid [][]byte) int {
	r := len(grid)
	c := len(grid[0])
	u := Init(grid)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i - 1 >= 0 && grid[i - 1][j] == '1' {
					u.Union((i - 1) * c + j, i * c + j)
				}
				if i + 1 < r && grid[i + 1][j] == '1' {
					u.Union((i + 1) * c + j, i * c + j)
				}
				if j - 1 >= 0 && grid[i][j - 1] == '1' {
					u.Union(i * c + j, i * c + j - 1)
				}
				if j + 1 < c && grid[i][j + 1] == '1' {
					u.Union(i * c + j + 1, i * c + j)
				}
			}
		}
	}

	return u.getGroups()
}


type unionFind struct {
	count int
	parents []int
	rank []int
}

func (this *unionFind) Find(x int) int {
	if this.parents[x] == x {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y int) {
	one := this.Find(x)
	two := this.Find(y)
	if one != two {
		this.parents[one] = two
		this.count = this.count - 1
	}
}

func Init(grid [][]byte) *unionFind {
	c := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				c++
			}
		}
	}
	u := &unionFind{count: c}
	u.parents = make([]int, m * n)
	for i := 0; i < m * n; i++ {
		u.parents[i] = i
	}
	return u
}

func (this *unionFind) getGroups() int {
	return this.count
}
