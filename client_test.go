package plivo

import (
	"context"
	"net/url"
	"os"
	"testing"
)

var authID = os.Getenv("AUTH_ID")
var authToken = os.Getenv("AUTH_TOKEN")
var fromNumber = os.Getenv("FROM_NUMBER")
var toNumber = os.Getenv("TO_NUMBER")

func TestMakeCall(t *testing.T) {
	c, err := NewClient(context.Background(), authID, authToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	u, err := url.Parse("https://s3.amazonaws.com/static.plivo.com/answer.xml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := c.MakeCall(fromNumber, toNumber, u)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
}

func TestGetCallDetails(t *testing.T) {
	c, err := NewClient(context.Background(), authID, authToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	res, err := c.GetCallDetails()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res == nil {
		t.Fatalf("unexpected error: %v", "response is empty")
	}
}
