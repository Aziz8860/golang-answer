package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func calculateDenseRanking(leaderboardScores []int, gitsScores []int) []int {
	uniqueSortedScores := removeDuplicatesAndSort(leaderboardScores)
	
	result := make([]int, len(gitsScores))
	for i, score := range gitsScores {
		rank := findRank(uniqueSortedScores, score)
		result[i] = rank
	}
	
	return result
}


func removeDuplicatesAndSort(scores []int) []int {
	scoreMap := make(map[int]bool)
	for _, score := range scores {
		scoreMap[score] = true
	}
	
	unique := make([]int, 0, len(scoreMap))
	for score := range scoreMap {
		unique = append(unique, score)
	}
	
	for i := 0; i < len(unique)-1; i++ {
		for j := i + 1; j < len(unique); j++ {
			if unique[i] < unique[j] {
				unique[i], unique[j] = unique[j], unique[i]
			}
		}
	}
	
	return unique
}

// findRank finds the rank of a given score in the uniqueSortedScores slice
func findRank(uniqueSortedScores []int, score int) int {
	for i, s := range uniqueSortedScores {
		if score >= s {
			return i + 1 // Ranks start from 1
		}
	}
	
	return len(uniqueSortedScores) + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter the number of players on the leaderboard:")
	scanner.Scan()
	nPlayers, _ := strconv.Atoi(scanner.Text())
	
	fmt.Println("Enter the leaderboard scores (space-separated):")
	scanner.Scan()
	scoresStr := strings.Fields(scanner.Text())
	leaderboardScores := make([]int, nPlayers)
	for i := 0; i < nPlayers && i < len(scoresStr); i++ {
		leaderboardScores[i], _ = strconv.Atoi(scoresStr[i])
	}
	
	fmt.Println("Enter the number of games GITS played:")
	scanner.Scan()
	nGames, _ := strconv.Atoi(scanner.Text())
	
	fmt.Println("Enter GITS's scores (space-separated):")
	scanner.Scan()
	gitsScoresStr := strings.Fields(scanner.Text())
	gitsScores := make([]int, nGames)
	for i := 0; i < nGames && i < len(gitsScoresStr); i++ {
		gitsScores[i], _ = strconv.Atoi(gitsScoresStr[i])
	}
	
	ranks := calculateDenseRanking(leaderboardScores, gitsScores)
	
	fmt.Println("\nOutput:")
	for i, rank := range ranks {
		fmt.Print(rank)
		if i < len(ranks)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
