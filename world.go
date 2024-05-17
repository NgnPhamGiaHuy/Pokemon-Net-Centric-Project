package main

type World struct {
	grid [][]int
}

func NewWorld() *World {
	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

	return &World{
		grid: grid,
	}
}
