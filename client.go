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
	userAgent = "go-plivo/0.1.0"
	baseURL   = "https://api.plivo.com"
)

type Client struct {
	ctx       context.Context
	baseURL   *url.URL
	authID    string
	authToken string
	*http.Client
}

func NewClient(ctx context.Context, authID, authToken string) (*Client, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{ctx, base, authID, authToken, http.DefaultClient}, nil
}

func (c *Client) Call(from, to string, answerURL *url.URL) (*http.Response, error) {
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

func (c *Client) newRequest(req *http.Request) (*http.Response, error) {
	req = req.WithContext(c.ctx)
	req.SetBasicAuth(c.authID, c.authToken)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}
