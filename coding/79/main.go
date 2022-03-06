package main

import "fmt"

func main()  {
	board := [][]byte{[]byte("ABCE"),[]byte("SFCS"),[]byte("ADEE")}
	fmt.Println(exist(board, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	visited := make([][]bool, row)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if search(board, word, i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, row, col, k int, visited [][]bool) bool {
	if board[row][col] == word[k] && !visited[row][col] {
		if k == len(word) - 1 {
			return true
		}
		visited[row][col] = true
		neighs := getNeighs(row, col, visited)
		for _, neigh := range neighs {
			if search(board, word, neigh[0], neigh[1], k + 1, visited) {
				return true
			}
		}
		visited[row][col] = false
		return false
	} else {
		return false
	}
}

func getNeighs(row, col int, visited [][]bool) [][]int {
	dirs := [][]int{{0,1},{1,0},{-1,0},{0,-1}}
	var neighs [][]int
	for _, dir := range dirs {
		newR := row + dir[0]
		newC := col + dir[1]
		if newR >= 0 && newR < len(visited) && newC >= 0 && newC < len(visited[0]) && !visited[newR][newC] {
			neighs = append(neighs, []int{newR, newC})
		}
	}
	return neighs
}
