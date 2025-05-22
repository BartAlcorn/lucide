package main

/**
Very Heavily inspired by Axel Adrian ( github: axzilla )

https://github.com/axzilla/goilerplate/blob/main/cmd/icongen/main.go

Generates icon definitions from SVG files in the `iconDir` directory
and saves them to `outputFile`/icon_defs.go

*/

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	iconDir = "./lucide/icons" // Path to the SVG files
	// outputFile     = "./icondefs.go"
	iconContentDir = "./icons" // Directory for individual icon contents
	packageName    = "icons"
)

func main() {
	// Read all files from the icon directory
	files, err := os.ReadDir(iconDir)
	if err != nil {
		panic(err)
	}

	// Initialize slice for icon definitions
	var iconDefs []string
	iconDefs = append(iconDefs, "package "+packageName+"\n")
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

			// Add icon definition
			// iconDefs = append(iconDefs, fmt.Sprintf("var %s = Icon(%q)\n", funcName, name))

			// Save icon content to a separate file
			content, err := os.ReadFile(filepath.Join(iconDir, file.Name()))
			if err != nil {
				panic(err)
			}

			var paths strings.Builder

			svgParts := strings.SplitN(string(content), ">", -1)
			count := len(svgParts) - 2
			for i := 1; i < count; i++ {
				paths.WriteString(fmt.Sprintf("  %s>", svgParts[i]))
			}
			content = []byte(fmt.Sprintf(TEMPL, funcName, paths.String()))

			err = os.WriteFile(filepath.Join(iconContentDir, name+".templ"), content, 0644)
			if err != nil {
				panic(err)
			}
		}
	}

	// Write all icon definitions to the output file
	// err = os.WriteFile(outputFile, []byte(strings.Join(iconDefs, "")), 0644)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Icon definitions and contents generated successfully!")
	templGenerate(iconContentDir)
}

// toPascalCase converts a kebab-case string to PascalCase
func toPascalCase(file string) string {
	var buf strings.Builder

	for _, sub := range strings.Split(file, "-") {
		if len(sub) < 1 {
			continue
		}
		buf.WriteString(strings.ToUpper(string(sub[0])))
		buf.WriteString(sub[1:])
	}

	return buf.String()
}

func templGenerate(target string) error {
	cmd := exec.Command("templ", "fmt", target)
	if _, err := cmd.Output(); err != nil {
		return err
	}

	cmd = exec.Command("templ", "generate")
	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}

const TEMPL = `package icons

templ %s(props ...Props) {
	{{ var p Props }}
    if len(props) > 0 {
  {{ p = props[0] }}
  }
	@SVG(p){
    %s
	}
}`
