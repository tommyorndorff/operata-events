package events

import "time"

// CallSummaryDetail represents the detail payload for CallSummary events
type CallSummaryDetail struct {
	AccountProperties AccountProperties `json:"accountProperties"`
	Contact           CallContact       `json:"contact"`
	WebRTCSession     WebRTCSession     `json:"webRTCSession"`
	ServiceAgent      ServiceAgent      `json:"serviceAgent"`
	Billing           Billing           `json:"billing"`
	Timestamp         time.Time         `json:"timestamp"`
}

// CallSummaryEvent represents a complete CallSummary EventBridge event
type CallSummaryEvent struct {
	EventBridgeEvent
	Detail CallSummaryDetail `json:"detail"`
}

// CallContact extends Contact with call-specific information
type CallContact struct {
	Contact
	Direction string     `json:"direction"`
	Events    CallEvents `json:"events"`
	EndedBy   string     `json:"endedBy"`
	QueueName string     `json:"queueName"`
	CallerID  string     `json:"callerId"`
}

// CallEvents represents call timing events
type CallEvents struct {
	ConnectingToAgent time.Time `json:"connectingToAgent"`
	Enqueued          time.Time `json:"enqueued"`
}

// WebRTCSession represents WebRTC session information
type WebRTCSession struct {
	Metrics            WebRTCMetrics      `json:"metrics"`
	ServiceEndpoint    ServiceEndpoint    `json:"serviceEndpoint"`
	MediaEndpoint      MediaEndpoint      `json:"mediaEndpoint"`
	SignallingEndpoint SignallingEndpoint `json:"signallingEndpoint"`
	UsedDevices        []Device           `json:"usedDevices"`
}

// WebRTCMetrics represents WebRTC quality metrics
type WebRTCMetrics struct {
	Inbound  InboundMetrics  `json:"inbound"`
	Outbound OutboundMetrics `json:"outbound"`
	RTT      RTTMetrics      `json:"rtt"`
	Jitter   JitterMetrics   `json:"jitter"`
	MOS      MOSMetrics      `json:"mos"`
}

// InboundMetrics represents inbound WebRTC metrics
type InboundMetrics struct {
	PacketsReceived       int          `json:"packetsReceived"`
	PacketsLost           int          `json:"packetsLost"`
	PacketsLostPercentage float64      `json:"packetsLostPercentage"`
	BytesReceived         int          `json:"bytesReceived"`
	AudioLevel            AudioLevel   `json:"audioLevel"`
	JitterBufferMils      JitterBuffer `json:"jitterBufferMils"`
}

// OutboundMetrics represents outbound WebRTC metrics
type OutboundMetrics struct {
	PacketsSent           int          `json:"packetsSent"`
	PacketsLost           int          `json:"packetsLost"`
	PacketsLostPercentage float64      `json:"packetsLostPercentage"`
	BytesSent             int          `json:"bytesSent"`
	AudioLevel            AudioLevel   `json:"audioLevel"`
	JitterBufferMils      JitterBuffer `json:"jitterBufferMils"`
}

// RTTMetrics represents round-trip time metrics
type RTTMetrics struct {
	Min int `json:"min"`
	Max int `json:"max"`
	Avg int `json:"avg"`
}

// JitterMetrics represents jitter metrics
type JitterMetrics struct {
	Min int `json:"min"`
	Max int `json:"max"`
	Avg int `json:"avg"`
}

// MOSMetrics represents Mean Opinion Score metrics
type MOSMetrics struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}

// ServiceEndpoint represents service endpoint information
type ServiceEndpoint struct {
	FQDN                     string    `json:"fqdn"`
	TransportLifeTimeSeconds int       `json:"transportLifeTimeSeconds"`
	Expiry                   time.Time `json:"expiry"`
}

// MediaEndpoint represents media endpoint information
type MediaEndpoint struct {
	FQDN            string `json:"fqdn"`
	DestinationPort string `json:"destinationPort"`
	SourcePort      string `json:"sourcePort"`
	Transport       string `json:"transport"`
	PrivateIP       string `json:"privateIp"`
}

// SignallingEndpoint represents signalling endpoint information
type SignallingEndpoint struct {
	FQDN string `json:"fqdn"`
}

// Device represents a device used during the call
type Device struct {
	Timestamp time.Time `json:"timestamp"`
	DeviceID  string    `json:"deviceId"`
	GroupID   string    `json:"groupId"`
	Kind      string    `json:"kind"`
	Label     string    `json:"label"`
}

// ServiceAgent represents agent information
type ServiceAgent struct {
	Username     string      `json:"username"`
	Machine      Machine     `json:"machine"`
	Network      Network     `json:"network"`
	Browser      Browser     `json:"browser"`
	Softphone    Softphone   `json:"softphone"`
	Interaction  Interaction `json:"interaction"`
	FriendlyName string      `json:"friendlyName"`
}

// Machine represents machine specifications and performance
type Machine struct {
	CPU    CPU    `json:"cpu"`
	Memory Memory `json:"memory"`
}

// CPU represents CPU information and metrics
type CPU struct {
	ModelName          string          `json:"modelName"`
	IdlePercentage     PercentageAvg   `json:"idlePercentage"`
	UtilisedPercentage PercentageRange `json:"utilisedPercentage"`
}

// PercentageAvg represents a percentage with average
type PercentageAvg struct {
	Avg float64 `json:"avg"`
}

// PercentageRange represents a percentage with min, max, and average
type PercentageRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}

// Memory represents memory information and metrics
type Memory struct {
	AvailableGB        float64         `json:"availableGb"`
	UtilisedPercentage PercentageRange `json:"utilisedPercentage"`
}

// Network represents network information
type Network struct {
	InternetGatewayIP string      `json:"internetGatewayIp"`
	MediaIPAddress    string      `json:"mediaIpAddress"`
	Type              string      `json:"type"`
	ISP               string      `json:"isp"`
	Geolocation       Geolocation `json:"geolocation"`
}

// Softphone represents softphone URLs
type Softphone struct {
	SoftphoneURL        string `json:"softphoneUrl"`
	SoftphoneContextURL string `json:"softphoneContextUrl"`
}

// Interaction represents interaction duration metrics
type Interaction struct {
	TotalDurationSec   int `json:"totalDurationSec"`
	OnHoldDurationSec  int `json:"onHoldDurationSec"`
	TalkingDurationSec int `json:"talkingDurationSec"`
	OnMuteDurationSec  int `json:"onMuteDurationSec"`
}

// Billing represents billing information
type Billing struct {
	DurationRoundedMin int `json:"durationRoundedMin"`
}
