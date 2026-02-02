package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strings"

	"github.com/goccy/go-yaml"
	"path/filepath"
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

func loadImport() []Style {
	resp, err := http.Get("https://github.com/catppuccin/userstyles/releases/download/all-userstyles-export/import.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	styleParse := []Style{}

	json.Unmarshal(data, &styleParse)

	return styleParse
}

func readBase16() map[string]string {
	user, _ := user.Current()

	data, _ := os.ReadFile("/home/" + user.Username + "/.config/stylix/palette.json")

	unm := map[string]string{}

	yaml.Unmarshal(data, &unm)

	return unm
}

func loadCustomStyles() []Style {
	out := []Style{}

	entries, err := os.ReadDir("customstyles")
	if err != nil {
		return out
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		ext := filepath.Ext(e.Name())
		if ext != ".css" && ext != ".scss" {
			continue
		}

		data, err := os.ReadFile(filepath.Join("customstyles", e.Name()))
		if err != nil {
			continue
		}

		out = append(out, Style{
			Enabled:    true,
			Name:       strings.TrimSuffix(e.Name(), ext),
			SourceCode: string(data),
		})
	}

	return out
}


func main() {
	styles := append(loadImport(), loadCustomStyles()...)

	modCol := []Style{}

	palette := readBase16()

	for _, s := range styles {
		if s.SourceCode == "" {
			continue
		}

		for col, key := range colors {
			if col == "accent" {
				reg := regexp.MustCompile(`\@accent(\)|\:|\!|\,|\ \=|\ \!)`)

				s.SourceCode = reg.ReplaceAllString(s.SourceCode, "#"+palette[key]+"$1")
			} else {
				s.SourceCode = strings.ReplaceAll(s.SourceCode, "@"+col, "#"+palette[key])
			}
		}
		modCol = append(modCol, s)
	}

	new, _ := json.Marshal(modCol)

	fmt.Println(string(new))
}
