package main

import (
	"math/rand"
	"time"
)

type Player struct {
	X, Y int
}

func NewPlayer(x, y int) *Player {
	return &Player{
		X: x,
		Y: y,
	}
}

func (p *Player) MoveUp() {
	if p.Y > 0 {
		p.Y--
	}
}

func (p *Player) MoveDown() {
	if p.Y < gridSize-1 {
		p.Y++
	}
}

func (p *Player) MoveLeft() {
	if p.X > 0 {
		p.X--
	}
}

func (p *Player) MoveRight() {
	if p.X < gridSize-1 {
		p.X++
	}
}

func (p *Player) AutoMove() {
	rand.Seed(time.Now().UnixNano())

	for {
		direction := rand.Intn(4)
		switch direction {
		case 0:
			p.MoveUp()
		case 1:
			p.MoveDown()
		case 2:
			p.MoveLeft()
		case 3:
			p.MoveRight()
		}
		time.Sleep(time.Second) // Adjust the sleep duration as needed
	}
}
