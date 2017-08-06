package plivo

type MakeCallOps struct {
	AnswerMethod           string `json:"answer_method,omitempty"`
	RingURL                string `json:"ring_url,omitempty"`
	RingMethod             string `json:"ring_method,omitempty"`
	HangupURL              string `json:"hangup_url,omitempty"`
	HangupMethod           string `json:"hangup_method,omitempty"`
	FallbackURL            string `json:"fallback_url,omitempty"`
	FallbackMethod         string `json:"fallback_method,omitempty"`
	CallerName             string `json:"caller_name,omitempty"`
	SendDigits             string `json:"send_digits,omitempty"`
	SendOnPreAnswer        bool   `json:"send_on_preanswer,omitempty"`
	TimeLimit              int    `json:"time_limit,omitempty"`
	HangupOnRing           int    `json:"hangup_on_ring,omitempty"`
	MachineDetection       bool   `json:"machine_detection,omitempty"`
	MachineDetectionTime   string `json:"machine_detection_time,omitempty"`
	MachineDetectionURL    string `json:"machine_detection_url,omitempty"`
	MachineDetectionMethod string `json:"machine_detection_method,omitempty"`
	SipHeaders             string `json:"sip_headers,omitempty"`
	RingTimeout            string `json:"ring_timeout,omitempty"`
	ParentCallUUID         string `json:"parent_call_uuid,omitempty"`
	ErrorIfParentNotFound  bool   `json:"error_if_parent_not_found,omitempty"`
}
