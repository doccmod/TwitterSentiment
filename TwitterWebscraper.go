package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	// other imports...

	twitterscraper "github.com/n0madic/twitter-scraper"
)

// SentimentAnalysis calls a Python script to analyze the sentiment of a given text.
func SentimentAnalysis(text string) (string, error) {
	cmd := exec.Command("python", "Tweet_sentiment_analysis.py", text)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return "", err
	}
	return result["sentiment"].(string), nil
}

func main() {
	scraper := twitterscraper.New()
	err := scraper.Login("__ajaffar", "Jk1231402")
	if err != nil {
		panic(err)
	}
	// Define your search query and the number of tweets you wish to analyze
	query := "AAPL OR Apple OR #AAPLStock"
	limit := 20

	// Create a CSV writer writing to stdout
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	err = writer.Write([]string{"Tweet", "Sentiment"})
	if err != nil {
		log.Fatalln("Error writing record to csv:", err)
	}

	for tweet := range scraper.SearchTweets(context.Background(), query, limit) {
		if tweet.Error != nil {
			log.Println("Error fetching tweet:", tweet.Error)
			continue
		}

		// Perform sentiment analysis on the tweet text
		sentiment, err := SentimentAnalysis(tweet.Text)
		if err != nil {
			fmt.Println("Error performing sentiment analysis:", err)
			continue // Skip this tweet if there's an error
		}

		// Write tweet text and sentiment to CSV
		err = writer.Write([]string{tweet.Text, sentiment})
		if err != nil {
			log.Fatalln("Error writing record to csv:", err)
		}
	}
}
