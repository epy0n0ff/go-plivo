package plivo

import (
	"context"
	"net/http/httputil"
	"net/url"
	"testing"
)

const (
	authID    = "your authID"
	authToken = "your authToken"
	from      = "11111111111"
	to        = "11111111111"
)

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
func TestCall(t *testing.T) {
	c, err := NewClient(context.Background(), authID, authToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	u, err := url.Parse("https://s3.amazonaws.com/static.plivo.com/answer.xml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := c.Call(from, to, u)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rawReq, err := httputil.DumpResponse(resp, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Logf("%s", string(rawReq))
}
