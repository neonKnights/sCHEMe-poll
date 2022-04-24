package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

var (
	layout           int
	max              int = 4
	windowW, windowH     = 1280, 960
)

var info struct {
	impression   string
	visualEffect string
	musicFeeling string
	playable     string
	plus         string
	minus        string
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Condition(layout == 4,
			giu.Layout{
				giu.Align(giu.AlignCenter).To(
					giu.Label("Dziękujemy za wypełnienie ankiety!"),
					giu.Button("Prześlij wynik").OnClick(onClick),
				),
			},
			giu.Layout{
				giu.Child().Layout(
					giu.Custom(func() {
						switch layout {
						case 0:
							mainLayout()
						case 1:
							firstPage()
						case 2:
							secondPage()
						case 3:
							thirdPage()
						}
					}),
				).Size(giu.Auto, float32(windowH-80)),
				giu.Row(
					giu.Button("<< Previous").Disabled(layout == 0).OnClick(func() {
						layout--
					}),
					giu.Button("Next >>").Disabled(layout == max).OnClick(func() {
						layout++
					}),
				),
			},
		),
	)
}

func mainLayout() {
	giu.Layout{
		giu.Align(giu.AlignCenter).To(
			giu.Label("Witaj w ankiecie na temat gry S(CHEM)E"),
			giu.Label("Jeśli masz chwilę, odpowiedz Nam na kilka pytań,"),
			giu.Label("Aby pomóc nam w rozwoju gry"),
		),
	}.Build()
}

func firstPage() {
	giu.Layout{
		giu.Label("Jakie jest twoje ogólne wrażenie z gry?"),
		giu.InputTextMultiline(&info.impression),
		giu.Separator(),
		giu.Label("Jakie są twoje odczucia co do warstwy wizualnej gry?"),
		giu.InputTextMultiline(&info.visualEffect),
	}.Build()
}

func secondPage() {
	giu.Layout{
		giu.Label("Oceń proszę warstwe muzyczną"),
		giu.InputTextMultiline(&info.musicFeeling),
		giu.Separator(),
		giu.Label("Czy gra jest grywalna?"),
		giu.InputTextMultiline(&info.playable),
	}.Build()
}

func thirdPage() {
	giu.Layout{
		giu.Label("Wymień zalety gry"),
		giu.InputTextMultiline(&info.plus),
		giu.Separator(),
		giu.Label("Wymień wady gry"),
		giu.InputTextMultiline(&info.minus),
	}.Build()
}

func onClick() {
	fmt.Println(info)
}

func main() {
	wnd := giu.NewMasterWindow("S(CHEM)E POLL", windowW, windowH, 0)
	wnd.Run(loop)
}
