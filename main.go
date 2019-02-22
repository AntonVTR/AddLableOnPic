package main

import (
	"log"

	"github.com/fogleman/gg"
)

func main() {

	//const S = 1024
	const S = 2484
	a := 2484
	b := 3512
	im, err := gg.LoadImage("src.jpg")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(a, b)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("Library/Fonts/Go-Mono.ttf", 96); err != nil {
		//if err := dc.LoadFontFace("Go-Mono.ttf", 96); err != nil {
		panic(err)
	}
	//dc.DrawStringAnchored("Hello, world!", S/3, S/3, 0.5, 0.5)

	//dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.Push()
	dc.DrawStringAnchored("Hello, world!", 1230, 1605, 0.5, 0.5)
	dc.DrawStringAnchored("300", 1381, 1771, 0.5, 0.5)
	dc.DrawStringAnchored("Старперы", 1430, 2117, 0.5, 0.5)
	dc.DrawStringAnchored("100:500", 400, 200, 0.5, 0.5)

	dc.Pop()
	dc.SavePNG("out.png")
}
