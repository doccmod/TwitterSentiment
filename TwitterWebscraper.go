package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

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
	limit := 10

	// Create a CSV writer writing to a file
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Updated to include "Date" in the header
	err = writer.Write([]string{"Tweet", "Sentiment", "Date"})
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

		// Format tweet's date
		tweetDate := tweet.TimeParsed.Format(time.RFC3339) // Ensure you have the correct property for tweet's date

		// Write tweet text, sentiment, and date to CSV
		err = writer.Write([]string{tweet.Text, sentiment, tweetDate})
		if err != nil {
			log.Fatalln("Error writing record to csv:", err)
		}
	}
}
