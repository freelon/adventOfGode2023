package main

import (
	"adventOfGode2023/day01"
	"adventOfGode2023/day02"
	"adventOfGode2023/day03"
	"adventOfGode2023/day04"
	"adventOfGode2023/day05"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Part = func(input string) string

var days = map[int][]Part{
	1: {day01.Part1, day01.Part2},
	2: {day02.Part1, day02.Part2},
	3: {day03.Part1, day03.Part2},
	4: {day04.Part1, day04.Part2},
	5: {day05.Part1, day05.Part2},
}

func main() {
	day := 5
	ensureInputExists(day)
	input := ReadFile(dailyInputPath(day))
	for part, f := range days[day] {
		start := time.Now()
		result := f(input)
		duration := time.Since(start)
		fmt.Printf("Day %d part %d result: %s\n", day, part+1, result)
		fmt.Printf("Took %s\n", duration)
	}
}

func ensureInputExists(day int) {
	var myFilePath = dailyInputPath(day)

	if doesFileExist(myFilePath) {
		fmt.Printf("file '%s' exists\n", myFilePath)
	} else {
		fmt.Printf("file '%s' does not exist. downloading ...\n", myFilePath)
		download(day)
	}
}

func download(day int) {
	var url = fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	var target = dailyInputPath(day)

	downloadUrlToFile(url, target)
}

func dailyInputPath(day int) string {
	return fmt.Sprintf("./input/day%02d.txt", day)
}

func doesFileExist(path string) (found bool) {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		} else {
			panic(fmt.Sprint("failed to check if file exists: ", err))
		}
	}
	return true
}

func downloadUrlToFile(fileUrl string, fileName string) {

	var cookie = strings.TrimSpace(ReadFile("cookie.env"))

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", fileUrl, nil)
	req.Header.Set("Cookie", cookie)

	// Put content on file
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)

}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
