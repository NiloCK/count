package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var repos []string = []string{
	"vue-skuilder",
	"tuido",
}
var startDate string = "2023"

func gitDiffReport() {
	allDifs := 0
	fmt.Println("GITDIFF REPORT:")
	for _, repo := range repos {
		allDifs += getDiffInRepo(repo)
	}
	fmt.Println("  -----------------")
	fmt.Printf("  Total: %d\n", allDifs)
}

func setWorkingDir(repo string) {
	os.Chdir("/home/colin/dev/" + repo)
}

func getFirstToken(s string) string {
	return strings.Split(s, " ")[0]
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getNumericTokens(s string) []int {
	tokens := strings.Split(s, " ")
	nums := []int{}
	for _, token := range tokens {
		if isNumeric(token) {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	return nums
}

func getDiffInRepo(repo string) int {
	setWorkingDir(repo)
	commitsBytes, err := exec.Command("git",
		"log",
		"--oneline",
		"--branches",
		"--after="+startDate,
		"--author=Colin").Output()
	commitsStr := string(commitsBytes)
	commits := strings.Split(commitsStr, "\n")

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(commits))

	commitHashes := []string{}
	for _, commit := range commits {
		commitHashes = append(commitHashes, getFirstToken(string(commit)))
	}

	diffs := [][]int{}
	for _, commit := range commitHashes {
		rawDiff, _ := exec.Command("git", "show", commit, "--oneline", "--shortstat").Output()
		// fmt.Println(string(rawDiff))
		diffs = append(diffs, getNumericTokens(string(rawDiff)))
	}

	totalDiff := 0
	for _, diff := range diffs {
		for i, num := range diff {
			// diff[0] is the number of files changed - do not count it
			if i != 0 {
				totalDiff += num
			}
		}
	}

	fmt.Printf("  %s: %d\n", repo, totalDiff)
	return totalDiff
}
