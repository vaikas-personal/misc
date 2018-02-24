package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tchap/go-patricia/patricia"
)

var allLetters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	letters := flag.String("letters", "", "letters to parse")
	length := flag.Int("length", 0, "how long of a word to create ..")

	flag.Parse()
	if *length == 0 {
		panic("Yo, need -length")
	}
	if len(*letters) == 0 {
		panic("Yo, need -letters")
	}

	file, err := os.Open("/home/vaikas/wordslist")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dict := make(map[string]int)
	trie := patricia.NewTrie()

	i := 0
	for scanner.Scan() { // internally, it advances token based on sperator
		//		fmt.Printf("%q\n", scanner.Text()) // token in unicode-char
		word := scanner.Text()
		trie.Insert(patricia.Prefix(word), i)
		dict[word] = 1
		i = i + 1
	}
	permute(*length, dict, *letters, "")
	validWords(*length, trie, *letters)
}

func permute(length int, words map[string]int, letters string, sofar string) {
	//	fmt.Printf("Getting called with: %q %q\n", letters, sofar)
	if len(letters) == 0 {
		return
	}
	for i, l := range letters {
		testString := sofar + string(l)
		if len(testString) == length {
			// If the letter is a * then run through all the possible letters it could
			// possibly be
			if l == '*' {
				for _, repl := range allLetters {
					testString2 := sofar + string(repl)
					if _, ok := words[testString2]; ok {
						fmt.Printf("FOUND word %q\n", testString2)
					}
				}
			} else {
				if _, ok := words[testString]; ok {
					fmt.Printf("FOUND word %q\n", testString)
				}
			}
		}
		//		fmt.Printf("CHECKING: %q\n", testString)
		remaining := letters[0:i] + letters[i+1:]
		if i < len(letters)-1 {
			permute(length, words, remaining, testString)
		}
	}
}

func validWords(length int, trie *patricia.Trie, prefix string) {
	printItem := func(prefix patricia.Prefix, item patricia.Item) error {
		if len(prefix) == length {
			fmt.Printf("%q: %v\n", prefix, item)
		} else {
			//fmt.Printf("****SKIPPING****%q: %v\n", prefix, item)
		}
		return nil
	}

	trie.VisitSubtree(patricia.Prefix(prefix), printItem)
}
