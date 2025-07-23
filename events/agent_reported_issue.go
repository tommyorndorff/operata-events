package events

import "time"

// AgentReportedIssueDetail represents the detail payload for AgentReportedIssue events
type AgentReportedIssueDetail struct {
	OperataClientID string         `json:"operataClientId"`
	Agent           string         `json:"agent"`
	State           string         `json:"state"`
	Context         IssueContext   `json:"context"`
	Browser         Browser        `json:"browser"`
	System          System         `json:"system"`
	SoftphoneError  SoftphoneError `json:"softphoneError"`
	Timestamp       time.Time      `json:"timestamp"`
	ID              string         `json:"id"`
}

// AgentReportedIssueEvent represents a complete AgentReportedIssue EventBridge event
type AgentReportedIssueEvent struct {
	EventBridgeEvent
	Detail AgentReportedIssueDetail `json:"detail"`
}

// IssueContext represents the context of a reported issue
type IssueContext struct {
	CallContactID string `json:"callContactId"`
	Category      string `json:"category"`
	Cause         string `json:"cause"`
	Message       string `json:"message"`
	Scenario      string `json:"scenario"`
	Severity      string `json:"severity"`
}

// System represents system information for issue reporting
type System struct {
	CPU    SystemCPU    `json:"cpu"`
	Memory SystemMemory `json:"memory"`
}

// SystemCPU represents CPU information in issue context
type SystemCPU struct {
	ModelName      string  `json:"modelName"`
	IdlePercentage float64 `json:"idlePercentage"`
	UsedPercentage float64 `json:"usedPercentage"`
}

// SystemMemory represents memory information in issue context
type SystemMemory struct {
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
}

// SoftphoneError represents softphone error information
type SoftphoneError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
