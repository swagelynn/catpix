package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

type Style struct {
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	UpdateURL   string `json:"updateUrl"`
	UsercssData struct {
		Name         string          `json:"name"`
		Namespace    string          `json:"namespace"`
		HomepageURL  string          `json:"homepageURL"`
		Version      string          `json:"version"`
		UpdateURL    string          `json:"updateURL"`
		SupportURL   string          `json:"supportURL"`
		Description  string          `json:"description"`
		Author       string          `json:"author"`
		License      string          `json:"license"`
		Preprocessor string          `json:"preprocessor"`
		Vars         json.RawMessage `json:"vars"`
	} `json:"usercssData"`
	SourceCode     string `json:"sourceCode"`
	OriginalDigest string `json:"originalDigest"`
}

var colors map[string]string = map[string]string{
	"crust":    "base00",
	"mantle":   "base00",
	"base":     "base00",
	"surface0": "base01",
	"surface1": "base04",
	"surface2": "base04",
	"overlay0": "base05",
	"overlay1": "base05",

	"overlay2": "base05",
	"text":     "base07",
	"subtext1": "base06",
	"subtext0": "base06",

	// pink
	"rosewater": "base08",
	"flamingo":  "base08",

	// purple
	"pink":     "base09",
	"mauve":    "base09",
	"lavender": "base09",

	// red-ish
	"red":    "base0A",
	"maroon": "base0A",
	"peach":  "base0A",

	// green/yellow/the cat threw up again
	"yellow": "base0B",
	"green":  "base0B",

	// bloo passport
	"teal": "base0C",
	"sky":  "base0C",

	// blue again i guess
	"sapphire": "base0D",
	"blue":     "base0D",

	"mocha":     "base0E",
	"latte":     "base0E",
	"frappe":    "base0E",
	"macchiato": "base0E",
	"accent":    "base0E",
}

// currently from a file, will implement fetching later
func loadImport() []Style {
	data, _ := os.ReadFile("/home/maddie/Downloads/import.json")

	styleParse := []Style{}

	json.Unmarshal(data, &styleParse)

	return styleParse
}

func readBase16() map[string]string {
	path := "/home/maddie/Downloads/gruvbox-dark-soft.yaml"

	data, _ := os.ReadFile(path)

	unm := struct {
		Palette map[string]string `yaml:"palette"`
	}{}

	yaml.Unmarshal(data, &unm)

	return unm.Palette
}

func main() {
	styles := loadImport()

	modCol := []Style{}

	palette := readBase16()

	for _, s := range styles {
		if s.SourceCode == "" {
			continue
		}

		for col, key := range colors {
			s.SourceCode = strings.ReplaceAll(s.SourceCode, "@"+col, palette[key])
		}

		modCol = append(modCol, s)
	}

	new, _ := json.Marshal(modCol)

	fmt.Println(string(new))
}
