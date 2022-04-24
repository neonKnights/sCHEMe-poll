package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"strings"

	"golang.org/x/oauth2"

	"github.com/AllenDang/giu"
	"github.com/google/go-github/v43/github"
)

//go:embed token.txt
var token string

var (
	layout           int
	max              int = 4
	windowW, windowH     = 1280, 960
	sent             bool
	dataname         string
)

type infoData struct {
	Impression   string `json:"impression"`
	VisualEffect string `json:"visualImpression"`
	MusicFeeling string `json:"musicFeeling"`
	Playable     string `json:"playable"`
	Plus         string `json:"pluses"`
	Minus        string `json:"minuses"`
}

var info infoData

func loop() {
	giu.SingleWindow().Layout(
		giu.Condition(layout == 4,
			giu.Layout{
				giu.Align(giu.AlignCenter).To(
					giu.Label("Dziękujemy za wypełnienie ankiety!"),
					giu.Button("Prześlij wynik").OnClick(onClick).Disabled(sent),
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
			giu.Row(
				giu.Label("Podaj swoje imię/nick lub co kolwiek innego po czym będziemy mogli cię zidentyfikować: "),
				giu.InputText(&dataname).Size(200),
			),
		),
	}.Build()
}

func firstPage() {
	giu.Layout{
		giu.Label("Jakie jest twoje ogólne wrażenie z gry?"),
		giu.InputTextMultiline(&info.Impression),
		giu.Separator(),
		giu.Label("Jakie są twoje odczucia co do warstwy wizualnej gry?"),
		giu.InputTextMultiline(&info.VisualEffect),
	}.Build()
}

func secondPage() {
	giu.Layout{
		giu.Label("Oceń proszę warstwe muzyczną"),
		giu.InputTextMultiline(&info.MusicFeeling),
		giu.Separator(),
		giu.Label("Czy gra jest grywalna?"),
		giu.InputTextMultiline(&info.Playable),
	}.Build()
}

func thirdPage() {
	giu.Layout{
		giu.Label("Wymień zalety gry"),
		giu.InputTextMultiline(&info.Plus),
		giu.Separator(),
		giu.Label("Wymień wady gry"),
		giu.InputTextMultiline(&info.Minus),
	}.Build()
}

func onClick() {
	if dataname == "" {
		log.Fatal("dataneme is empty!")
	}

	issueText := dataname + "said:\n```json\n"
	data, err := json.MarshalIndent(info, "\n", "\n")
	if err != nil {
		log.Fatal("error encoding json: %w")
	}

	issueText += string(data)
	issueText += "\n```"

	issueComment := &github.IssueComment{
		Body: &issueText,
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	_, _, err = client.Issues.CreateComment(ctx, "neonKnights", "sCHEMe-poll", 1, issueComment)
	if err != nil {
		log.Fatalf("Errore createing issue: %v", err)
	}

	sent = true
}

func main() {
	token = strings.ReplaceAll(token, "\n", "")
	wnd := giu.NewMasterWindow("S(CHEM)E POLL", windowW, windowH, 0)
	wnd.Run(loop)
}
