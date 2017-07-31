package plivo

import "time"

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+timeFormat+`"`, string(data))
	return err
}

type CallResult struct {
	ApiID       string `json:"api_id"`
	Message     string `json:"message"`
	RequestUUID string `json:"request_uuid"`
}

type CallDetails struct {
	ApiID   string   `json:"api_id"`
	Meta    Meta     `json:"meta"`
	Objects []Detail `json:"objects"`
}

type Meta struct {
	Limit      int `json:"limit"`
	Next       int `json:"next"`
	Offset     int `json:"offset"`
	Previous   int `json:"previous"`
	TotalCount int `json:"total_count"`
}

type Detail struct {
	AnswerTime     *Time  `json:"answer_time"`
	BillDuration   int    `json:"bill_duration"`
	BilledDuration int    `json:"billed_duration"`
	CallDirection  string `json:"call_direction"`
	CallDuration   int    `json:"call_duration"`
	CallUUID       string `json:"call_uuid"`
	EndTime        *Time  `json:"end_time"`
	FromNumber     string `json:"from_number"`
	InitiationTime *Time  `json:"initiation_time"`
	ParentCallUUID string `json:"parent_call_uuid"`
	ResourceURI    string `json:"resource_uri"`
	ToNumber       string `json:"to_number"`
	TotalAmount    string `json:"total_amount"`
	TotalRate      string `json:"total_rate"`
}
