package events

import "time"

// HeartbeatWorkflowEvent represents a heartbeat workflow event item
type HeartbeatWorkflowEvent struct {
	AgentID          string    `json:"agentId"`
	AgentType        string    `json:"agentType"`
	AxScore          int       `json:"axScore"`
	CreatedOn        time.Time `json:"createdOn"`
	CxScore          int       `json:"cxScore"`
	DiallerCallID    string    `json:"diallerCallId"`
	GroupID          string    `json:"groupId"`
	HeartbeatID      string    `json:"heartbeatId"`
	JobID            string    `json:"jobId"`
	NetworkScore     int       `json:"networkScore"`
	ReceiverCallID   string    `json:"receiverCallId"`
	RoutingProfileID string    `json:"routingProfileId"`
	Status           string    `json:"status"`
	ToNumber         string    `json:"toNumber"`
	UpdatedOn        time.Time `json:"updatedOn"`
}

// HeartbeatWorkflowEvents represents an array of heartbeat workflow events
type HeartbeatWorkflowEvents []HeartbeatWorkflowEvent
