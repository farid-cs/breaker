package main

import rl "github.com/gen2brain/raylib-go/raylib"

const WINDOW_WIDTH = 1600
const WINDOW_HEIGHT = 900

type Rectangle struct {
	width int32
	height int32
	x int32
	y int32
	color rl.Color
}

func main() {
	plate := Rectangle {
		200,
		60,
		WINDOW_WIDTH / 2 - 100,
		WINDOW_HEIGHT * 3 / 4 - 40,
		rl.Red,
	}

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "pingpog")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
	        if rl.IsKeyDown(rl.KeyH) {
			plate.x -= 5
		}
		if rl.IsKeyDown(rl.KeyL) {
			plate.x += 5
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			plate.x -= 5
		}
		if rl.IsKeyDown(rl.KeyRight) {
			plate.x += 5
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(plate.x, plate.y, plate.width, plate.height, plate.color)

		rl.EndDrawing()
	}
}
