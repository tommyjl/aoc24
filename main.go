package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type AoC struct{}

func main() {
	day := os.Getenv("DAY")
	t := reflect.TypeFor[AoC]()
	pad := ""

	for i := range t.NumMethod() {
		m := t.Method(i)

		if day != "" {
			ok, _ := regexp.MatchString(`^SolveDay`+day+`$`, m.Name)
			if !ok {
				continue
			}
		}

		fmt.Printf(pad)
		pad = "\n"
		fmt.Println(m.Name)
		fmt.Println(strings.Repeat("=", len(m.Name)))

		inputs := []reflect.Value{
			reflect.ValueOf(AoC{}),
		}
		m.Func.Call(inputs)
	}
}
