package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// keep legacy name to avoid breaking nip rules
var canonicalNameOverridesByID = map[int]string{
	27:  "Sabre",            // was generated as Saber
	41:  "Kris",             // was generated as Kriss
	135: "Stiletto",         // was generated as Stilleto
	166: "LargeSiegeBow",    // was generated as LongSiegeBow
	236: "MithrilPoint",     // was generated as MithralPoint
	379: "AncientShield",    // was generated as KurastShield (xts)
	417: "DemonHeadShield",  // was generated as DemonHead (necro head)
	428: "DemonHead",        // was generated as Demonhead (helm)
	157: "QuarterStaff",     // was generated as Quarterstaff
	275: "DemonCrossBow",    // was generated as DemonCrossbow
	469: "GriffonHeaddress", // was generated as GriffonHeadress
}

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

	nameIDRe := regexp.MustCompile(`Name:\s*"([^"]*)".*\bID:\s*([0-9]+)\b`)
	namesByID := make(map[int]string)
	seenCanonicalRaw := make(map[string]string)
	maxID := -1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		m := nameIDRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		raw := m[1]
		id, err := strconv.Atoi(m[2])
		if err != nil {
			return fmt.Errorf("parsing item id %q: %w", m[2], err)
		}
		canonical := canonicalizeName(raw)
		if canonical == "" {
			return fmt.Errorf("unable to generate canonical name for item %q", raw)
		}
		if override, ok := canonicalNameOverridesByID[id]; ok {
			canonical = override
		}

		if existing, exists := namesByID[id]; exists && existing != canonical {
			return fmt.Errorf("item id %d produced inconsistent canonical names %q and %q", id, existing, canonical)
		}
		namesByID[id] = canonical
		key := strings.ToLower(canonical)
		if existingRaw, exists := seenCanonicalRaw[key]; exists && existingRaw != raw {
			return fmt.Errorf("duplicate generated item name %q from %q and %q", canonical, existingRaw, raw)
		}
		if _, exists := seenCanonicalRaw[key]; !exists {
			seenCanonicalRaw[key] = raw
		}
		if id > maxID {
			maxID = id
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if maxID < 0 {
		return fmt.Errorf("no item names found in pkg/data/item/items.go")
	}

	names := make([]string, maxID+1)
	for id := 0; id <= maxID; id++ {
		name, ok := namesByID[id]
		if !ok {
			return fmt.Errorf("missing item name for id %d in pkg/data/item/items.go", id)
		}
		names[id] = name
	}

	fmt.Printf("Extracted %d names from items.go\n", len(names))

	// Write name.go
	var b bytes.Buffer
	b.WriteString(`package item

import "strings"

const (
	ScrollOfTownPortal = "ScrollOfTownPortal"
	ScrollOfIdentify   = "ScrollOfIdentify"
	TomeOfTownPortal   = "TomeOfTownPortal"
	TomeOfIdentify     = "TomeOfIdentify"
	Key                = "Key"
)

func GetNameByEnum(itemNumber uint) Name {
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
