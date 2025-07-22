package events

// HeadsetSummaryDetail represents the detail payload for HeadsetSummary events
type HeadsetSummaryDetail struct {
	AccountProperties AccountProperties `json:"accountProperties"`
	Contact           HeadsetContact    `json:"contact"`
	Headset           Headset           `json:"headset"`
}

// HeadsetSummaryEvent represents a complete HeadsetSummary EventBridge event
type HeadsetSummaryEvent struct {
	EventBridgeEvent
	Detail HeadsetSummaryDetail `json:"detail"`
}

// HeadsetContact extends Contact with headset-specific interaction data
type HeadsetContact struct {
	Contact
	Interaction HeadsetInteraction `json:"interaction"`
	QueueName   string             `json:"queueName"`
}

// HeadsetInteraction represents interaction data for headset events
type HeadsetInteraction struct {
	TotalDurationSec            int `json:"totalDurationSec"`
	OnHoldDurationSec           int `json:"onHoldDurationSec"`
	AgentInteractionDurationSec int `json:"agentInteractionDurationSec"`
}

// Headset represents headset device information and metrics
type Headset struct {
	ModelName       string         `json:"modelName"`
	FirmwareVersion string         `json:"firmwareVersion"`
	SerialNumber    string         `json:"serialNumber"`
	APIVersion      string         `json:"apiVersion"`
	Metrics         HeadsetMetrics `json:"metrics"`
}

// HeadsetMetrics represents metrics collected from headset
type HeadsetMetrics struct {
	Speech                  SpeechMetrics `json:"speech"`
	ExposureDB              DBMetrics     `json:"exposureDb"`
	BackgroundNoiseDB       DBMetrics     `json:"backgroundNoiseDb"`
	MisalignedBoomArmCount  int           `json:"misalignedBoomArmCount"`
	DeviceMuteCount         int           `json:"deviceMuteCount"`
	DeviceVolumeAdjustCount int           `json:"deviceVolumeAdjustCount"`
}

// SpeechMetrics represents speech-related metrics from headset
type SpeechMetrics struct {
	CrossTalkTotal    float64 `json:"crossTalkTotal"`
	CrossTalkTotalPct float64 `json:"crossTalkTotalPct"`
	RxSpeechTotal     float64 `json:"rxSpeechTotal"`
	RxSpeechTotalPct  float64 `json:"rxSpeechTotalPct"`
	SilenceTotal      float64 `json:"silenceTotal"`
	SilenceTotalPct   float64 `json:"silenceTotalPct"`
	TotalSeconds      float64 `json:"totalSeconds"`
	TxSpeechTotal     float64 `json:"txSpeechTotal"`
	TxSpeechTotalPct  float64 `json:"txSpeechTotalPct"`
}

// DBMetrics represents decibel-related metrics
type DBMetrics struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}
