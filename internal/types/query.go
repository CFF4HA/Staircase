package types

import (
	"github.com/google/uuid"
)

// We have three types of actions: click, extract-text, and
// input (keyboard input a value into a field).
type Action struct {
	// click, extract-text, input
	Type  string `json:"type"`
	Value string `json:"value,omitempty"`
}

type Navigation struct {
	Keyword    string `json:"keyword"`
	Quantifier int    `json:"quantifier"`
}

type RetrievalQuery struct {
	Site        string       `json:"site"`
	Regex       string       `json:"regex"`
	Navigations []Navigation `json:"navigations"`
	Action      Action       `json:"action"`

	// The following is meta-information required
	// for tracking the query via the database.
	Id uuid.UUID `json:"id"`
}

// This is the overall structure that we will
// pass and it will contain all query objects, etc.
type RetrievalJob struct {
	Queries []RetrievalQuery `json:"queries"`

	Datasource *DatabaseDatasource `json:"datasource"`
	Id         uuid.UUID           `json:"id"`
}
