package files

import (
	"bufio"
	"os"
	"unicode"
)

func WordFreq(file string) map[string]int {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	//var words []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var dict = make(map[string]int)

	for scanner.Scan() {
		_len := len(scanner.Text())
		if !unicode.IsPunct(rune(scanner.Text()[_len - 1])) {
			dict[scanner.Text()]++
		} else {
			dict[string([]byte(scanner.Text())[:_len - 1])]++
		}
	}

	return dict
}
