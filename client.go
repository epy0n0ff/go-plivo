package plivo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	userAgent  = "go-plivo/0.2.0"
	baseURL    = "https://api.plivo.com"
	timeFormat = "2006-01-02 15:04:05Z07:00"
)

type Client struct {
	baseURL   *url.URL
	authID    string
	authToken string
	*http.Client
}

func NewClient(authID, authToken string) (*Client, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{base, authID, authToken, http.DefaultClient}, nil
}

func (c *Client) MakeCall(ctx context.Context, from, to string, answerURL *url.URL, ops *MakeCallOps) (*CallResult, error) {
	u := c.baseURL
	u.Path = fmt.Sprintf("/v1/Account/%s/Call/", c.authID)

	body := struct {
		To        string `json:"to"`
		From      string `json:"from"`
		AnswerURL string `json:"answer_url"`
		*MakeCallOps
	}{
		To:          to,
		From:        from,
		AnswerURL:   answerURL.String(),
		MakeCallOps: ops,
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

	res, err := c.newRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("invalid status code:%d", res.StatusCode)
	}

	var result CallResult
	if err := c.decodeBody(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCallDetails(ctx context.Context) (*CallDetails, error) {
	u := c.baseURL
	u.Path = fmt.Sprintf("/v1/Account/%s/Call/", c.authID)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.newRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code:%d", res.StatusCode)
	}

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

func (c *Client) newRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	req.SetBasicAuth(c.authID, c.authToken)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}
