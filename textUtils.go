package main

import (
    "strings"
    "unicode"
    "unicode/utf8"
    "strconv"
    _"fmt"
    "regexp"
)

func isShort(word string, minSize int) bool {
    return utf8.RuneCountInString(word) < minSize
}

func isIn(word string, words []string) bool {
    for _, w := range words {
        if strings.ToLower(w) == strings.ToLower(word) {
            return true
        }
    }

    return false
}

func removePunctuation(text string) string {
    // match a sequence of unicode letters which may contain an apostrophe
    re := regexp.MustCompile("\\p{L}+('\\p{L}+|\\p{L}*)")
    return strings.Join(re.FindAllString(text, -1), " ")
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
    mapCopy := make(map[string]int)
    for k, v := range wordCount {
        mapCopy[k] = v
    }

    for word1, _ := range mapCopy {
        for word2, _ := range wordCount {
            jwDistance := Calculate(word1, word2)
            if word1 != word2 && jwDistance > similarity {
                // fmt.Println(word1, "\t", word2)
                delete(wordCount, word2)
                delete(mapCopy, word2)
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
        result.WriteString(k + " ")
    }

    return result.String()
}

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