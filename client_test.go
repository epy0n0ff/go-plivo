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

/*
HTTP/1.1 201 CREATED
Content-Length: 139
Connection: keep-alive
Content-Type: application/json
Date: Sun, 09 Jul 2017 18:33:37 GMT
Server: nginx/1.4.6 (Ubuntu)
X-Plivo-Signature: xxxxx

{
	"api_id": "xxxx",
 	 "message": "call fired",
 	 "request_uuid": "xxxxxxx"
}
*/
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
