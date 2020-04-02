package main

import (
    "strings"
    "unicode"
    "unicode/utf8"
)

func removeUpperCase(text string) string {
    return strings.ToLower(text)
}

func removeChars(text string, chars []rune) string {
    for _, r := range chars {
        text = strings.ReplaceAll(text, string(r), " ")
    }

    return text
}

func isShort(word string, minSize int) bool {
    if utf8.RuneCountInString(word) >= minSize {
        return false
    } else {
        return true
    }
}

func isIn(word string, words []string) bool {
    for _, w := range words {
        if strings.ToLower(w) == strings.ToLower(word) {
            return true
        }
    }

    return false
}

func removeShortWords(text string, minSize int) string {
    goodLengthWords := []string{}

    words := strings.Split(text, " ")
    for _, word := range words {
        if utf8.RuneCountInString(word) >= minSize {
            goodLengthWords = append(goodLengthWords, word)
        }
    }

    return strings.Join(goodLengthWords, " ")
}

func getSpecialChars(text string) []rune {
    specialChars := make(map[rune]bool)
    specialCharsSlice := []rune{}

    for _, r := range text {
        // if a special char is not already stored
        if !unicode.IsLetter(r) && !specialChars[r] {
            specialChars[r] = true
            specialCharsSlice = append(specialCharsSlice, r)
        }
    }

    return specialCharsSlice
}

func filterWordCounter(wordCount map[string]int, minWordSize int, wordsToRemove []string) {
    for word, _ := range wordCount {
        if isShort(word, minWordSize) || isIn(word, wordsToRemove) {
            delete(wordCount, word)
        }
    }
}

func getWordCount(text string) map[string]int {
    words := strings.Fields(text)

    wordCounter := make(map[string]int)
    for _, word := range words {
        if wordCounter[word] > 0 {
            wordCounter[word]++
        } else {
            wordCounter[word] = 1
        }
    }

    return wordCounter
}

func removeWords(words map[string]int, wordsToRemove []string) {
    for _, word := range wordsToRemove {
        delete(words, word)
    }
}

// remove punctuation first
func getCapitalizedWords(text string) []string {
    capitalizedWords := []string{}

    words := strings.Fields(text)
    for _, word := range words {
        if unicode.IsUpper([]rune(word)[0]) {
            capitalizedWords = append(capitalizedWords, word)
        } 
    }

    return capitalizedWords
}

func suggestProperNouns(capitalizedWords []string, minOccurences int) []string {
    // const minOccurences int = 5;

    suggestions := []string{}

    text := strings.Join(capitalizedWords, " ")
    wordCounter := getWordCount(text)

    for word, count := range wordCounter {
        if count >= minOccurences {
            suggestions = append(suggestions, word)
        }
    }

    return suggestions
}