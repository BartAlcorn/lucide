package lucide

/*
	Very Heavily inspired by Axel Adrian ( github: axzilla )
	https://github.com/axzilla/goilerplate/blob/main/cmd/icongen/main.go

	The major difference being we pre-extract and only save the internal svg data, not the entire .svg file.
	This allows us to save/store less data (about 200 bytes per icon) and more importantly, it is one less step
	during processing/display, which happens many, many more times than the initial file import.

	* download https://github.com/lucide-icons/lucide separately
	* then create a symlink to the lucide directory
	* for example, from within the ./cmd/genlucide folder:
	! ln -s  ../../../icons/lucide/icons icons

	* then run from the project root folder
	! go run ./cmd/genlucide

*/

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LucideVersion represents the version of Lucide icons used in this package.
// Here for reference only.
const LucideVersion = "0.462.0"

const (
	// * this is an symlink/alias to the actual icon directory
	iconDir        = "./cmd/genlucide/icons"           // Path to the Lucide SVG files
	outputFile     = "./pkg/icons/lucide/icon_defs.go" // Output file for icon definitions
	iconContentDir = "./pkg/icons/lucide/content"      // Directory for individual icon contents
)

func main() {
	// Read all files from the icon directory
	files, err := os.ReadDir(iconDir)
	if err != nil {
		panic(err)
	}

	// Initialize slice for icon definitions
	var iconDefs []string
	iconDefs = append(iconDefs, "package lucide\n")
	iconDefs = append(iconDefs, "// This file is auto generated\n")

	// Create the content directory if it doesn't exist
	err = os.MkdirAll(iconContentDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Process each SVG file
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".svg" {
			name := strings.TrimSuffix(file.Name(), ".svg")
			funcName := toPascalCase(name)

			// Add to the icon definition slice
			iconDefs = append(iconDefs, fmt.Sprintf("var %s = Icon(%q)\n", funcName, name))

			// read the original SVG file
			content, err := os.ReadFile(filepath.Join(iconDir, file.Name()))
			if err != nil {
				panic(err)
			}

			// extrace the only the actual SVG path data
			// we do this now so we're storing the smallest amount of data/bytes
			// and we remove one extra step during processing and display
			extractedContemt := extractSVGContent(string(content))
			err = os.WriteFile(filepath.Join(iconContentDir, name+".svg"), []byte(extractedContemt), 0644)
			if err != nil {
				panic(err)
			}

		}
	}

	// Write all icon definitions to the output file
	err = os.WriteFile(outputFile, []byte(strings.Join(iconDefs, "")), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Icon definitions and contents generated successfully!")
}

// toPascalCase converts a kebab-case string to PascalCase
func toPascalCase(s string) string {
	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// extractSVGContent removes the outer SVG tags from the icon content.
func extractSVGContent(svgContent string) string {
	start := strings.Index(svgContent, ">") + 1
	end := strings.LastIndex(svgContent, "</svg>")
	if start == -1 || end == -1 {
		return ""
	}
	return strings.TrimSpace(svgContent[start:end])
}
