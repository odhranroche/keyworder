/*
 * Description: This program is intended to extract the minimum number of words
 *              needed to understand the text. Duplicate words, similar words,
 *              short words, and proper nouns are removed. The resulting list 
 *              forms the foundational words of the text.
 *
 * Usage: go build main.go textUtils.go jaro.go
 *        .\main.go input.txt
 *
 * Issues: - Word similarity algorithm sometimes deletes words
             with different meanings
           - Suggested proper nouns sometimes picks words that
             are at the start of a sentence
 */

package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "time"
)

func main() {
    start := time.Now()

    // number of times a word is capitalized to count it as a proper noun
    const maxProperNounOccurrences int = 3

    // words shorter than this will be removed
    const minWordSize int = 4
    
    // words more similar than this will be removed
    const maxSimilarity float64 = 0.9;

    if len(os.Args) != 2 {
        fmt.Println("Usage:", os.Args[0], "file")
        return
    }
    var inputFile string = string(os.Args[1])
    var outputFile string = inputFile + "_output.txt"

    text := fileToString(inputFile)
    text = removePunctuation(text)    

    totalWords := len(strings.Fields(text))
    fmt.Println("Total words: ", totalWords)

    // count occurances of words & remove duplicates
    wordCounter := getWordCount(text)

    // take a list of all likely proper nouns 
    capitalizedWords := getCapitalizedWords(text)
    suggestedNouns := suggestProperNouns(capitalizedWords, maxProperNounOccurrences)

    // remove short words and proper nouns
    filterWordCounterByWords(wordCounter, suggestedNouns)
    filterWordCounterBySize(wordCounter, minWordSize)
    filterWordCounterBySimilarity(wordCounter, maxSimilarity)

    fmt.Println("Final size: ", len(wordCounter))
    saveToFile(outputFile, mapKeysToString(wordCounter))

    elapsed := time.Since(start)
    fmt.Printf("Time: %s", elapsed)
}

func fileToString(filename string) string {
    fileContentBytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    
    return string(fileContentBytes)
}

func saveToFile(filename string, text string) error {
    return ioutil.WriteFile(filename, []byte(text), 0666)
}