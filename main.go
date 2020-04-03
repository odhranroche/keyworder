/*
 * Description: This program is intended to extract the minimum number of words
 *              needed to understand the text. Duplicate words, similar words,
 *              short words, and proper nouns are removed. The resulting list 
 *              forms the foundational words of the text.
 *
 * Usage: go run main.go textUtils.go jaro.go
 *        - input file should be in the same directory 
 *        - input file should be specified in main.go
 *        - number of times a word is capitalized to count as a proper noun
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
    const maxSimilarity float64 = 0.95;

    // const inputFile string = "Госпожа и господин Дърсли, живеещи на улица „Привит Драйв“"
    // const inputFile string = "Mister and misses Dursley, who lived on 'Privet Drive'"
    // const inputFile string = "HP_Short.txt"
    const inputFile string = "HP1_Eng.txt"
    // const inputFile string = "HP_Full.txt"
    const outputFile string = inputFile + "_output"

    text := fileToString(inputFile)
    
    // remove punctuation
    specialChars := getSpecialChars(text)
    text = removeChars(text, specialChars)

    totalWords := len(strings.Fields(text))
    fmt.Println("Total words: ", totalWords)

    // take a list of all likely proper nouns 
    capitalizedWords := getCapitalizedWords(text)
    suggestedNouns := suggestProperNouns(capitalizedWords, maxProperNounOccurrences)

    // count occurances of words
    wordCounter := getWordCount(text)

    // remove short words and proper nouns
    filterWordCounterBySize(wordCounter, minWordSize)
    filterWordCounterByWords(wordCounter, suggestedNouns)
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