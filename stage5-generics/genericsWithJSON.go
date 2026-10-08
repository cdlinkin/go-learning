/*
exercise:
Working with files and JSON
1. The program reads a large text file (≥100 MB, generate it yourself) line by line using bufio, counts the word frequency, and outputs the top 20. Measure the time and memory consumption.
2. Reads a JSON file with an array of objects, filters by a field, and writes the result to another JSON with indentation.
3. Handle the case of unknown fields using Decoder.DisallowUnknownFields()
*/
package stage5generics

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

func countWords(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)

	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		for _, word := range words {
			word = strings.ToLower(word)
			word = strings.Trim(word, ".,!?;:\"()[]{}")

			if word != "" {
				wordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordCount, nil
}

type WordFrequency struct {
	Word  string
	Count int
}

func printTop20(wordCount map[string]int) {
	words := make([]WordFrequency, 0, len(wordCount))

	for word, count := range wordCount {
		words = append(words, WordFrequency{
			Word:  word,
			Count: count,
		})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	if len(words) > 20 {
		words = words[:20]
	}

	fmt.Println("top 20 words:")

	for i, word := range words {
		fmt.Printf("%d. %-20s %d\n", i+1, word.Word, word.Count)
	}
}

func filterUsers(inputFile, outputFile string) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	decoder := json.NewDecoder(inFile)

	decoder.DisallowUnknownFields()

	var users []User

	if err := decoder.Decode(&users); err != nil {
		return fmt.Errorf("decode JSON: %w", err)
	}

	var filtered []User

	for _, user := range users {
		if user.Role == "backend" {
			filtered = append(filtered, user)
		}
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)

	encoder.SetIndent("", "  ")

	if err := encoder.Encode(filtered); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}

func RunWithJSON() {
	fmt.Println("1. Word frequency")

	start := time.Now()

	wordCount, err := countWords("input.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("processing time: %v\n", time.Since(start))

	printTop20(wordCount)

	fmt.Println()
	fmt.Println("2. JSON filtering")

	if err := filterUsers("input.json", "filtered.json"); err != nil {
		fmt.Println("error:", err)
		return
	}
}
