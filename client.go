package plivo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	userAgent  = "go-plivo/0.1.0"
	baseURL    = "https://api.plivo.com"
	timeFormat = "2006-01-02 15:04:05Z07:00"
)

type Client struct {
	ctx       context.Context
	baseURL   *url.URL
	authID    string
	authToken string
	*http.Client
}

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

func NewClient(ctx context.Context, authID, authToken string) (*Client, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{ctx, base, authID, authToken, http.DefaultClient}, nil
}

func (c *Client) MakeCall(from, to string, answerURL *url.URL) (*http.Response, error) {
	u := c.baseURL
	u.Path = fmt.Sprintf("/v1/Account/%s/Call/", c.authID)

	body := struct {
		To           string `json:"to"`
		From         string `json:"from"`
		AnswerURL    string `json:"answer_url"`
		AnswerMethod string `json:"answer_method"`
	}{
		To:           to,
		From:         from,
		AnswerURL:    answerURL.String(),
		AnswerMethod: "GET",
	}

	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return nil, err
	}

	return c.newRequest(req)
}

func (c *Client) GetCallDetails() (*CallDetails, error) {
	u := c.baseURL
	u.Path = fmt.Sprintf("/v1/Account/%s/Call/", c.authID)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.newRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var details CallDetails
	if err := c.decodeBody(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

func (c *Client) decodeBody(res *http.Response, v interface{}) error {
	d := json.NewDecoder(res.Body)
	return d.Decode(v)
}

func (c *Client) newRequest(req *http.Request) (*http.Response, error) {
	req = req.WithContext(c.ctx)
	req.SetBasicAuth(c.authID, c.authToken)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}
