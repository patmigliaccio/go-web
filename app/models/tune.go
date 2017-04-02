package models

// Tune : data model for a tune object
type Tune struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
