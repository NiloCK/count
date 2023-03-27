package main

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strings"
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

type QualifiedFile struct {
	dir  string
	file fs.FileInfo
}

func getMDFilesInDirRecursive(dir string) []QualifiedFile {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	retFiles := []QualifiedFile{}
	for _, file := range files {
		if file.IsDir() {
			subFiles := getMDFilesInDirRecursive(dir + "/" + file.Name())
			retFiles = append(retFiles, subFiles...)
		} else {
			if strings.HasSuffix(file.Name(), ".md") {
				retFiles = append(retFiles, QualifiedFile{dir, file})
			}
		}
	}
	return retFiles
}

func getFilesInBlogDir() []QualifiedFile {
	return getMDFilesInDirRecursive("/home/colin/blog")
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

func getFileWordCount(file QualifiedFile) int {
	if file.file.IsDir() {
		return 0
	}

	f, err := os.Open(file.dir + "/" + file.file.Name())
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
