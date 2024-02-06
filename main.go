package main

func main() {
	win := Window{}
	win.CreateWindow(Vector2{800, 600}, "My Win")
	defer win.DeleteWindow()

	for win.running() {
		win.StartDraw()

		win.EndDraw()
	}
}
