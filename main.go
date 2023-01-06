package main

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"time"
)

var goalWPD int = 100
var initWords int = 1563

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	blogReport()
	gitDiffReport()
}

func getFilesInBlogDir() []fs.FileInfo {
	dir, err := os.Open("/home/colin/blog")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func daysSinceLastYear() int {
	now := time.Now()
	newYears := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)

	days := now.Sub(newYears).Hours() / 24
	return int(days)
}

// totalWordCount returns the total number of
// words in all files in the blog directory
func totalWordCount() int {
	count := 0
	files := getFilesInBlogDir()

	for _, file := range files {
		count += getFileWordCount(file)
	}
	return count
}

func getFileWordCount(file fs.FileInfo) int {
	if file.IsDir() {
		return 0
	}

	f, err := os.Open("/home/colin/blog/" + file.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wordCount
}
