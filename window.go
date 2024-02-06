package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct{}

func (win Window) CreateWindow(size Vector2, name string) {
	rl.InitWindow(size.x, size.y, name)

	rl.SetTargetFPS(60)
}

func (win Window) DeleteWindow() {
	rl.CloseWindow()
}

func (win Window) running() bool {
	return !rl.WindowShouldClose()
}

func (win Window) StartDraw() {
	rl.BeginDrawing()
}

func (win Window) EndDraw() {
	rl.EndDrawing()
}
