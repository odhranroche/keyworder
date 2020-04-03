package main

import (
    "strings"
    "unicode"
    "unicode/utf8"
    "strconv"
)

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

func removeChars(text string, chars []rune) string {
    for _, r := range chars {
        text = strings.ReplaceAll(text, string(r), " ")
    }

    return text
}

func getWordCount(text string) map[string]int {
    words := strings.Fields(text)

    wordCounter := make(map[string]int)
    for _, word := range words {
        word = strings.ToLower(word)
        if wordCounter[word] > 0 {
            wordCounter[word]++
        } else {
            wordCounter[word] = 1
        }
    }

    return wordCounter
}

func filterWordCounterBySize(wordCount map[string]int, minWordSize int) {
    for word, _ := range wordCount {
        if isShort(word, minWordSize) {
            delete(wordCount, word)
        }
    }
}

func filterWordCounterByWords(wordCount map[string]int, wordsToRemove []string) {
    for word, _ := range wordCount {
        if isIn(word, wordsToRemove) {
            delete(wordCount, word)
        }
    }
}

func filterWordCounterBySimilarity(wordCount map[string]int, similarity float64) {
    keys := mapKeysToSlice(wordCount)

    for _, word1 := range keys {
        for word2, _ := range wordCount {
            jwDistance := Calculate(word1, word2)
            if word1 != word2 && jwDistance > similarity {
                delete(wordCount, word2)
            }
        }
    }    
}

func mapToString(m map[string]int) string {
    var result strings.Builder
    for k, v := range m {
        result.WriteString(k + "," + strconv.Itoa(v) + "\n")
    }

    return result.String()
}

func mapKeysToString(m map[string]int) string {
    var result strings.Builder
    for k, _ := range m {
        result.WriteString(k + "\n")
    }

    return result.String()
}

func mapKeysToSlice(m map[string]int) []string {
    result := []string{}
    for k, _ := range m {
        result = append(result, k)
    }

    return result
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