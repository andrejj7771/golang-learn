package main

import (
	"fmt"
	"github.com/andrejj7771/golang-learn/files"
)

func main() {
	//f, err := ioutil.ReadFile("data/test_data/files/wordfreq.txt")
	dict := files.WordFreq("data/test_data/files/wordfreq.txt")
	fmt.Println(dict)
}
