package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func stringFunctions() {
	p("Contains ", strings.Contains("test", "es"))
	p("Count ", strings.Count("test", "t"))
	p("HasPrefix ", strings.HasPrefix("test", "te"))
	p("HasSuffix ", strings.HasSuffix("test", "st"))
	p("Index ", strings.Index("test", "e"))
	p("Joins ", strings.Join([]string{"te", "s"}, "-"))
	p("Repeat ", strings.Repeat("test", 5))
	p("Replace ", strings.Replace("test", "t", "&", 1))
	p("ReplaceAll ", strings.ReplaceAll("test", "t", "&"))
	p("Split ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower ", strings.ToLower("TEST"))
	p("ToUpper ", strings.ToUpper("test"))
}
