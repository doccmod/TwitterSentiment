# sentiment_analysis.py
from textblob import TextBlob
import sys
import json

def analyze_sentiment(text):
    analysis = TextBlob(text)
    return analysis.sentiment.polarity

if __name__ == '__main__':
    text = sys.argv[1]
    sentiment_score = analyze_sentiment(text)
    result = {"sentiment": "Neutral"}
    if sentiment_score > 0:
        result["sentiment"] = "Positive"
    elif sentiment_score < 0:
        result["sentiment"] = "Negative"
    print(json.dumps(result))
