package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const defTab string = "  "

func mkTabs(n int) (out string) {
	for i := 0; i <= n; i++ {
		out += defTab
	}
	return
}

func printH1(w1, w2 string) {
	s1 := color.YellowString("%s:", strings.ToUpper(w1))
	s2 := fmt.Sprintf("'%s'", w2)
	fmt.Printf("%s %s\n", s1, s2)
}

func printH2(w string, n int) {
	s1 := color.YellowString("%s:", strings.ToUpper(w))
	s2 := fmt.Sprintf("(%d)", n)
	fmt.Printf("\n%s %s\n", s1, s2)
}

func printDef1(n int, def, pos string) {
	s1 := color.YellowString("%4d.", n)
	s2 := color.MagentaString(" %9s ", pos)
	fmt.Printf("%s %s %s\n", s1, s2, def)
}
