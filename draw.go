package main

import (
	"github.com/fogleman/gg"
)

func main() {
	drawContext := gg.NewContext(1000, 1000)
	drawContext.DrawCircle(500, 500, 400)
	drawContext.SetRGB(0, 0, 0)
	drawContext.Fill()
	drawContext.SavePNG("out.png")
}