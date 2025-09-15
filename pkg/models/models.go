package models


import "encoding/json"


type IngestItem struct {
SourceID string `json:"source_id"`
Kind string `json:"kind"` // tweet, article, macro, price
Data json.RawMessage `json:"data"`
}


type Article struct {
ID int64 `json:"id"`
Title string `json:"title"`
Text string `json:"text"`
SourceID int64 `json:"source_id"`
PublishedAt string `json:"published_at"`
URL string `json:"url"`
Sentiment float32 `json:"sentiment"`
}


type MacroEvent struct {
ID int64 `json:"id"`
Region string `json:"region"`
Name string `json:"name"`
ReleaseTS string `json:"release_ts"`
Consensus float64 `json:"consensus"`
Previous float64 `json:"previous"`
Actual *float64 `json:"actual,omitempty"`
Surprise *float64 `json:"surprise,omitempty"`
}