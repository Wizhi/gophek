package gophek

import (
	"io"
)

func Eval(source string, w io.Writer) {
	mem := make([]rune, 30000)

	var p int

	for i := 0; i < len(source); i++ {
		switch source[i] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			mem[p]++
		case '-':
			mem[p]--
		case '[':
			if mem[p] == 0 {
				for i < len(source) {
					i++

					if source[i] == ']' {
						break
					}
				}
			}
		case ']':
			if mem[p] == 0 {
				continue
			}

			for {
				i--

				if source[i] == '[' {
					break
				}
			}
		case '.':
			w.Write([]byte(string(mem[p])))
		case ',':
		default:
		}
	}
}
