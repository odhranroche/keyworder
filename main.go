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

    const maxProperNounOccurrences int = 3
    const minWordSize int = 4
    const maxSimilarity float64 = 0.95;

    // const inputFile string = "Госпожа и господин Дърсли, живеещи на улица „Привит Драйв“"
    // const inputFile string = "Mister and misses Dursley, who lived on 'Privet Drive'"
    // const inputFile string = "HP_Short.txt"
    const inputFile string = "HP_1.txt"
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