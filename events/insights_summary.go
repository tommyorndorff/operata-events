package events

// InsightsSummaryDetail represents the detail payload for InsightsSummary events
type InsightsSummaryDetail struct {
	AccountProperties AccountProperties `json:"accountProperties"`
	Contact           Contact           `json:"contact"`
	Insights          Insights          `json:"insights"`
}

// InsightsSummaryEvent represents a complete InsightsSummary EventBridge event
type InsightsSummaryEvent struct {
	EventBridgeEvent
	Detail InsightsSummaryDetail `json:"detail"`
}

// Insights represents insights data with tags
type Insights struct {
	Count int          `json:"count"`
	Tags  []InsightTag `json:"tags"`
}

// InsightTag represents an individual insight tag
type InsightTag struct {
	Description string `json:"description"`
}
