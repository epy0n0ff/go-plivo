package plivo

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"testing"
)

var authID = os.Getenv("AUTH_ID")
var authToken = os.Getenv("AUTH_TOKEN")
var fromNumber = os.Getenv("FROM_NUMBER")
var toNumber = os.Getenv("TO_NUMBER")
var callBackURL = os.Getenv("CALLBACK_URL")

func TestMakeCall(t *testing.T) {
	c, err := NewClient(authID, authToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	u, err := url.Parse(fmt.Sprintf("%s%s", callBackURL, "/plivo/answer"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ops := &MakeCallOps{
		AnswerMethod:   "POST",
		RingURL:        fmt.Sprintf("%s%s", callBackURL, "/plivo/ring"),
		RingMethod:     "POST",
		HangupURL:      fmt.Sprintf("%s%s", callBackURL, "/plivo/hangup"),
		HangupMethod:   "POST",
		FallbackURL:    fmt.Sprintf("%s%s", callBackURL, "/plivo/fallback"),
		FallbackMethod: "POST",
	}
	ctx := context.Background()
	res, err := c.MakeCall(ctx, fromNumber, toNumber, u, ops)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res == nil {
		t.Fatalf("unexpected error: %v", "response is empty")
	}
}

func TestGetCallDetails(t *testing.T) {
	c, err := NewClient(authID, authToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ctx := context.Background()
	res, err := c.GetCallDetails(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res == nil {
		t.Fatalf("unexpected error: %v", "response is empty")
	}
}
