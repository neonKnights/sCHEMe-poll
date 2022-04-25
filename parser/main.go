package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type infoData struct {
	Impression   string `json:"impression"`
	VisualEffect string `json:"visualImpression"`
	MusicFeeling string `json:"musicFeeling"`
	Playable     string `json:"playable"`
	Plus         string `json:"pluses"`
	Minus        string `json:"minuses"`
}

func main() {
	dataRecords := make([]infoData, 0)
input:
	for {
		var dataStr string
		var out infoData
		fmt.Print("Enter valid json input: ")
		for {
			c := make([]byte, 1)
			_, err := os.Stdin.Read(c)
			if err != nil {
				log.Fatal("error reading from stdin " + err.Error())
			}

			if c[0] == '$' {
				break input
			}
			dataStr += string(c)
			if c[0] == '}' {
				break
			}
		}

		err := json.Unmarshal([]byte(dataStr), &out)
		if err != nil {
			log.Fatalf("error detected (%v) starting processing!", err)
		}

		dataRecords = append(dataRecords, out)
	}

	out, err := gocsv.MarshalBytes(dataRecords)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(dataRecords))
	fmt.Println(string(out))
	os.WriteFile("output.csv", out, 0o644)
}
