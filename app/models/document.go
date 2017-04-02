package models

// Document : data model for a document object
type Document struct {
	ID         string `json:"id"`
	RecordType string `json:"record_type"`
	Value      string `json:"value"`
}
