package events

import "time"

// EventBridgeEvent represents the common structure for all EventBridge events
type EventBridgeEvent struct {
	Version    string      `json:"version"`
	ID         string      `json:"id"`
	DetailType string      `json:"detail-type"`
	Source     string      `json:"source"`
	Account    string      `json:"account"`
	Time       time.Time   `json:"time"`
	Region     string      `json:"region"`
	Resources  []string    `json:"resources"`
	Detail     interface{} `json:"detail"`
}

// AccountProperties represents common account information
type AccountProperties struct {
	OperataGroupName string `json:"operataGroupName,omitempty"`
	OperataGroupID   string `json:"operataGroupId"`
}

// ContactID represents contact identification structure
type ContactID struct {
	Current  string `json:"current"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
}

// Contact represents contact information common across events
type Contact struct {
	ID ContactID `json:"id"`
}

// AudioLevel represents audio level metrics
type AudioLevel struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}

// JitterBuffer represents jitter buffer metrics
type JitterBuffer struct {
	Min int     `json:"min"`
	Max int     `json:"max"`
	Avg float64 `json:"avg"`
}

// Browser represents browser information
type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Geolocation represents geographical location data
type Geolocation struct {
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
}
