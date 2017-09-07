package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tchap/go-patricia/patricia"
)

func main() {
	letters := flag.String("letters", "", "letters to parse")
	length := flag.Int("length", 0, "a string")

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

	trie := patricia.NewTrie()

	i := 0
	for scanner.Scan() { // internally, it advances token based on sperator
		//		fmt.Printf("%q\n", scanner.Text()) // token in unicode-char
		trie.Insert(patricia.Prefix(scanner.Text()), i)
		i = i + 1
	}
	permute(*length, *letters, "")
	validWords(*length, trie, *letters)
}

func permute(length int, letters string, sofar string) {
	fmt.Printf("Getting called with: %q %q\n", letters, sofar)
	if len(letters) == 0 {
		return
	}
	for i, l := range letters {
		testString := sofar + string(l)
		fmt.Printf("CHECKING: %q\n", testString)
		remaining := letters[0:i] + letters[i+1:]
		if i < len(letters)-1 {
			permute(length, remaining, testString)
		}
	}
}

func validWords(length int, trie *patricia.Trie, letters string) {
	fmt.Printf("YO: %s\n", letters)
	printItem := func(prefix patricia.Prefix, item patricia.Item) error {
		if len(prefix) == length {
			fmt.Printf("%q: %v\n", prefix, item)
		} else {
			fmt.Printf("****SKIPPING****%q: %v\n", prefix, item)
		}
		return nil
	}

	trie.VisitSubtree(patricia.Prefix(letters), printItem)
}
