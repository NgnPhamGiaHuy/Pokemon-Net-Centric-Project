package main

import (
	"fmt"
)

const gridSize = 1000

func main() {
	world := NewWorld()

	player := NewPlayer(gridSize/2, gridSize/2)
	go player.AutoMove()

	for {
		var input string
		fmt.Print("Enter direction (up, down, left, right): ")
		fmt.Scanln(&input)

		switch input {
		case "up":
			player.MoveUp()
		case "down":
			player.MoveDown()
		case "left":
			player.MoveLeft()
		case "right":
			player.MoveRight()
		default:
			fmt.Println("Invalid direction.")
		}

		fmt.Printf("Player position: (%d, %d)\n", player.X, player.Y)
	}
}
