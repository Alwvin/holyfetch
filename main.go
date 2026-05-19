package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

type Gospel struct {
	Id         int
	Version    string
	Verse_Name string
	Verse_Text string
}

func main() {
	resp, err := http.Get("your_bible.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("API request failed with status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var matius_tb []Gospel
	err = json.Unmarshal(body, &matius_tb)
	if err != nil {
		panic(err)
	}

	random := rand.Intn(len(matius_tb))

	ascii := `
        ██
        ██
		██
 ████████████████
        ██
        ██
        ██
		██
		██
		██
		██
		██
`

	header := matius_tb[random].Verse_Name + " " + matius_tb[random].Version

	cleanText := matius_tb[random].Verse_Text
	cleanText = strings.ReplaceAll(cleanText, "\t", " ")
	cleanText = strings.ReplaceAll(cleanText, "\r", "")
	cleanText = strings.ReplaceAll(cleanText, "\n", " ")
	cleanText = strings.ReplaceAll(cleanText, "⠀", "")
	cleanText = strings.TrimSpace(cleanText)

	asciiLines := strings.Split(ascii, "\n")
	for i := range asciiLines {
		asciiLines[i] = strings.ReplaceAll(asciiLines[i], "\t", "    ")
	}

	rightWidth := 40
	wrappedVerse := wrapText(cleanText, rightWidth)

	textLines := []string{header, ""}
	textLines = append(textLines, wrappedVerse...)

	maxLines := max(len(asciiLines), len(textLines))
	topPadding := (len(textLines) - len(asciiLines)) / 2
	if topPadding < 0 {
		topPadding = 0
	}

	leftWidth := 28

	for i := 0; i < maxLines; i++ {
		left := ""
		right := ""

		asciiIndex := i - topPadding
		if asciiIndex >= 0 && asciiIndex < len(asciiLines) {
			left = asciiLines[asciiIndex]
		}

		if i < len(textLines) {
			right = textLines[i]
		}

		paddedLeft := fmt.Sprintf("%-*s", leftWidth, left)
		fmt.Printf("%s | %s\n", paddedLeft, right)
	}
}

func wrapText(text string, width int) []string {
	words := strings.Fields(text)
	var lines []string
	current := ""

	for _, word := range words {
		if len(current)+len(word)+1 > width {
			lines = append(lines, current)
			current = word
		} else {
			if current != "" {
				current += " "
			}
			current += word
		}
	}

	if current != "" {
		lines = append(lines, current)
	}

	return lines
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
