package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    const maxProperNounOccurrences int = 3
    const minWordSize int = 4

    // text := "Госпожа и господин Дърсли, живеещи на улица „Привит Драйв“"
    text := fileToString("HP_1.txt")
    // text := fileToString("HP_Short.txt")
    
    // remove punctuation
    specialChars := getSpecialChars(text)
    text = removeChars(text, specialChars)

    // take a list of all likely proper nouns 
    capitalizedWords := getCapitalizedWords(text)
    suggestedNouns := suggestProperNouns(capitalizedWords, maxProperNounOccurrences)

    // lower case
    text = strings.ToLower(text)

    // count occurances of words
    wordCounter := getWordCount(text)

    // remove short words and proper nouns
    filterWordCounter(wordCounter, minWordSize, suggestedNouns)

    saveToFile("HP_output", mapToString(wordCounter))
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