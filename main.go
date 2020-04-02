package main

import (
    "fmt"
    "io/ioutil"
    "os"
    //"strings"
)

func main() {

    const maxProperNounOccurrences int = 2
    const minWordSize int = 3
    // text := "Госпожа и господин Дърсли, живеещи на улица „Привит Драйв“"
    // text := fileToString("HP_1.txt")
    text := fileToString("HP_Short.txt")
    
    // 1 remove special chars
    specialChars := getSpecialChars(text)
    text = removeChars(text, specialChars)

    // 2 take a list of all likely proper nouns 
    capitalizedWords := getCapitalizedWords(text)
    suggestedNouns := suggestProperNouns(capitalizedWords, maxProperNounOccurrences)

    // 3 lower case
    text = removeUpperCase(text)

    // fmt.Println(text)
    
    wordCount := getWordCount(text)

    fmt.Println("Before filtering:")
    fmt.Println(wordCount)
    fmt.Println(len(wordCount))
    
    filterWordCounter(wordCount, minWordSize, suggestedNouns)

    fmt.Println("After filtering:")
    fmt.Println(wordCount)
    fmt.Println(len(wordCount))
    // 4 remove short words
    //text = removeShortWords(text, 4)


    //removeWords(wordCount, suggestedNouns)

    //result := strings.Join(capitalizedWords, " ")
    //pop := getWordPopularity(result)
    //fmt.Println(pop)
    //for k, v := range words {
    //    fmt.Println(k, v)
    //}

//    fmt.Println(popularity)
//    fmt.Println(len(popularity))
}

func fileToString(filename string) string {
    fileContentBytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    
    return string(fileContentBytes)
}