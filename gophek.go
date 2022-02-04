package gophek

import (
	"bufio"
	"io"
)

func Eval(source string, r io.Reader, w io.Writer) {
	mem := make([]rune, 30000)
	inp := bufio.NewReader(r)

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
			r, n, err := inp.ReadRune()

			if err != nil {
				panic(err)
			}

			if n == 0 {
				return
			}

			mem[p] = r
		default:
		}
	}
}
