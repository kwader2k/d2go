package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// Read items.go
	f, err := os.Open("pkg/data/item/items.go")
	if err != nil {
		return err
	}
	defer f.Close()

	nameRe := regexp.MustCompile(`Name:\s*"([^"]*)"`)
	var names []string
	seenNames := make(map[string]string)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		m := nameRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		raw := m[1]
		canonical := canonicalizeName(raw)
		if canonical == "" {
			return fmt.Errorf("unable to generate canonical name for item %q", raw)
		}

		if existingRaw, exists := seenNames[canonical]; exists && existingRaw != raw {
			return fmt.Errorf("duplicate generated item name %q from %q and %q", canonical, existingRaw, raw)
		}
		if _, exists := seenNames[canonical]; !exists {
			seenNames[canonical] = raw
		}
		names = append(names, canonical)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("Extracted %d names from items.go\n", len(names))

	// Write name.go
	var b bytes.Buffer
	b.WriteString(`package item

import "strings"

const (
	ExpCharItemIDThreshold = 508
	ExpCharItemIDOffset    = 15
)

var ExpChar uint16

func SetExpChar(expChar uint16) {
	ExpChar = expChar
}

const (
	ScrollOfTownPortal = "ScrollOfTownPortal"
	ScrollOfIdentify   = "ScrollOfIdentify"
	TomeOfTownPortal   = "TomeOfTownPortal"
	TomeOfIdentify     = "TomeOfIdentify"
	Key                = "Key"
)

func GetNameByEnum(itemNumber uint) Name {
	idx := int(itemNumber)
	if ExpChar >= 3 && idx >= ExpCharItemIDThreshold {
		idx += ExpCharItemIDOffset
	}
	if idx < 0 || idx >= len(Names) {
		return Name("")
	}
	return Name(Names[idx])
}

func GetIDByName(itemName string) int {
	for i, name := range Names {
		if strings.EqualFold(name, itemName) {
			if ExpChar >= 3 && i >= ExpCharItemIDThreshold+ExpCharItemIDOffset {
				return i - ExpCharItemIDOffset
			}
			return i
		}
	}

	return -1
}

type Name string

var Names = []string{
`)
	for _, n := range names {
		fmt.Fprintf(&b, "\t\"%s\",\n", n)
	}
	b.WriteString("}\n")

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("formatting generated name.go: %w", err)
	}

	if err := os.WriteFile("pkg/data/item/name.go", formatted, 0644); err != nil {
		return err
	}

	fmt.Println("Generated name.go successfully")
	return nil
}

func canonicalizeName(raw string) string {
	raw = strings.TrimSpace(raw)
	raw = strings.ReplaceAll(raw, "'", "")

	var canonical strings.Builder
	capitalizeNext := true

	for _, r := range raw {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if capitalizeNext && unicode.IsLetter(r) {
				r = unicode.ToUpper(r)
			}
			canonical.WriteRune(r)
			capitalizeNext = false
			continue
		}

		capitalizeNext = true
	}

	return canonical.String()
}
