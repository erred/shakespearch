package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("shakespeare.json")
	defer f.Close()
	f2, _ := os.Create("shakespeare_v1.json")
	defer f2.Close()
	f2.WriteString("[\n")
	s := bufio.NewScanner(f)
	var x bool
	for s.Scan() {
		if x {
			f2.WriteString(",\n")
		}
		f2.Write(s.Bytes())
		x = true
	}
	f2.WriteString("]")
}
