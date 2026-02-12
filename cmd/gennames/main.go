package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	// Read items.go
	f, err := os.Open("pkg/data/item/items.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nameRe := regexp.MustCompile(`Name:\s*"([^"]*)"`)
	var names []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		m := nameRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		raw := strings.TrimSpace(m[1])
		// Remove apostrophes
		raw = strings.ReplaceAll(raw, "'", "")
		// Split by whitespace, capitalize each word, join
		words := strings.Fields(raw)
		var camel strings.Builder
		for _, w := range words {
			if len(w) == 0 {
				continue
			}
			runes := []rune(w)
			runes[0] = unicode.ToUpper(runes[0])
			camel.WriteString(string(runes))
		}
		names = append(names, camel.String())
	}

	fmt.Printf("Extracted %d names from items.go\n", len(names))

	// Write name.go
	out, err := os.Create("pkg/data/item/name.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	w.WriteString(`package item

import "strings"

const (
	ScrollOfTownPortal = "ScrollOfTownPortal"
	ScrollOfIdentify   = "ScrollOfIdentify"
	TomeOfTownPortal   = "TomeOfTownPortal"
	TomeOfIdentify     = "TomeOfIdentify"
	Key                = "Key"
)

func GetNameByEnum(itemNumber uint) Name {
	if int(itemNumber) >= len(Names) {
		return Name("")
	}
	return Name(Names[itemNumber])
}

func GetIDByName(itemName string) int {
	for i, name := range Names {
		if strings.EqualFold(name, itemName) {
			return i
		}
	}

	return -1
}

type Name string

var Names = []string{
`)
	for _, n := range names {
		fmt.Fprintf(w, "\t\"%s\",\n", n)
	}
	w.WriteString("}\n")
	w.Flush()

	fmt.Println("Generated name.go successfully")
}
