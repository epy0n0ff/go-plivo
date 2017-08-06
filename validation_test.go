package plivo

import (
	"net/url"
	"testing"
)

func TestValidateSignature(t *testing.T) {
	signature := "2+dMGSbMzxXZfAAkN4e8Hf+NmrQ="
	uri := "https://example.com/ring"
	param := "Direction=outbound&From=%2B819000000000&To=818000000000&RequestUUID=ad7106c9-8ac5-4c4f-93d4-df17bb62eee8&CallUUID=5b4a0ab4-be76-4e64-9e3b-f66dcde87ee2&CallStatus=ringing&Event=Ring"
	authToken := "YzU2N2RlYTYtN2E3Ni0xMWU3LWJiMzEtYmUyZTQ0YjA2YjM0"
	err := ValidateSignature(uri, param, signature, authToken)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSortedValues(t *testing.T) {
	param := "Direction=outbound&From=%2B819000000000&To=818000000000&RequestUUID=ad7106c9-8ac5-4c4f-93d4-df17bb62eee8&CallUUID=5b4a0ab4-be76-4e64-9e3b-f66dcde87ee2&CallStatus=ringing&Event=Ring"
	expected := "CallStatusringingCallUUID5b4a0ab4-be76-4e64-9e3b-f66dcde87ee2DirectionoutboundEventRingFrom+819000000000RequestUUIDad7106c9-8ac5-4c4f-93d4-df17bb62eee8To818000000000"
	v, _ := url.ParseQuery(param)
	actual := sortedValues(v).String()

	if expected != actual {
		t.Fatalf("unexpected error: %s!=%s", expected, actual)
	}
}
