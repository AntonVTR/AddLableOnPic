package Draw

import (
	"log"

	"github.com/fogleman/gg"
)

func AddText(name string, tName string, lTime string, place string) {

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
	dc.DrawImage(im, 0, 0)
	dc.Push()
	dc.DrawStringAnchored(name, 1230, 1605, 0.5, 0.5)
	dc.DrawStringAnchored(place, 1381, 1771, 0.5, 0.5)
	dc.DrawStringAnchored(lTime, 350, 200, 0.5, 0.5)

	dc.Pop()
	if err := dc.LoadFontFace("Library/Fonts/Go-Mono.ttf", 66); err != nil {
		//if err := dc.LoadFontFace("Go-Mono.ttf", 96); err != nil {
		panic(err)
	}
	dc.Push()

	dc.DrawStringAnchored(tName, 1430, 2117, 0.5, 0.5)
	dc.Pop()

	dc.SavePNG("pic/" + name + ".png")
}
