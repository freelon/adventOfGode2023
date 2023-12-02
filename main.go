package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

type Part = func(input string) string

var days = map[int][]Part{
	1: {Day1Part1, Day1Part2},
	2: {Day2Part1, Day2Part2},
}

func main() {
	day := 2
	ensureInputExists(day)
	input := ReadFile(dailyInputPath(day))
	//input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	for part, f := range days[day] {
		result := f(input)
		fmt.Printf("Day %d part %d result: %s\n", day, part, result)
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
