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

func main() {
	ensureInputExists(1)
	input := ReadFile(dailyInputPath(1))
	result := Day1Part1(input)
	fmt.Printf("Day 1 part 1 result: %d\n", result)
	result = Day1Part2(input)
	fmt.Printf("Day 1 part 2 result: %d", result)
}

func ensureInputExists(day int8) {
	var myFilePath = dailyInputPath(day)

	if doesFileExist(myFilePath) {
		fmt.Printf("file '%s' exists\n", myFilePath)
	} else {
		fmt.Printf("file '%s' does not exist. downloading ...\n", myFilePath)
		download(day)
	}
}

func download(day int8) {
	var url = fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	var target = dailyInputPath(day)

	downloadUrlToFile(url, target)
}

func dailyInputPath(day int8) string {
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

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
