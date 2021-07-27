package main

import (
	"fmt"
	"strings"
	"math"
)

var replaceChars = []string{",", ":", ".", "!", "?", ".", ";"}
const replacementChar = " "

func calcDistance(doc1 map[string]int, doc2 map[string]int) float64 {
	docProduct := 0
	for word, count1  := range doc1  {
		if count2, exist := doc2[word]; exist {
			docProduct += count1 * count2
		} 
	}
	docsLenProduct := len(doc1) * len(doc2)

	fmt.Printf("%v, %v\n", docProduct, docsLenProduct)
	
	return math.Acos(float64(docProduct)/float64(docsLenProduct))
}

func calcWords(doc string) map[string]int {
	for _, ch := range replaceChars {
		doc = strings.Replace(doc, ch, replacementChar, 0)
	}
	docWords := strings.Split(doc, replacementChar)
	words := make(map[string]int)
	for _, word := range docWords {
		if count, exist := words[word]; exist {
			words[word] = count + 1
		} else {
			words[word] = 1
		}
	}
	return words
}

func compareDocs(doc1 string, doc2 string) {
        distance := calcDistance(calcWords(doc1), calcWords(doc2))
	fmt.Printf("distance between %s, %s is %.2f\n", doc1, doc2, distance)
}

func main() {
	compareDocs("the dog is my pet", "the cat is my pet")
	compareDocs("the dog", "the cat")
	compareDocs("dog", "cat")
	compareDocs("the", "the")
}
