package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func SentimentAnalysis(text string) (string, error) {
	cmd := exec.Command("python", "sentiment_analysis.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	var result map[string]string
	if err := json.Unmarshal(output, &result); err != nil {
		return "", err
	}
	return result["sentiment"], nil
}

func main() {
	scraper := twitterscraper.New()
	err := scraper.Login("__ajaffar", "Jk1231402")
	if err != nil {
		panic(err)
	}

	for tweet := range scraper.SearchTweets(context.Background(), "tesla", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}

		// Perform sentiment analysis on the tweet text
		sentiment, err := SentimentAnalysis(tweet.Text)
		if err != nil {
			fmt.Println("Error performing sentiment analysis:", err)
			continue // Skip this tweet if there's an error
		}

		// Print the tweet text along with its sentiment
		fmt.Printf("Tweet: %s\nSentiment: %s\n\n", tweet.Text, sentiment)
	}
}
