package plivo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
)

type sortedValues url.Values

func (s sortedValues) String() string {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	b := bytes.NewBufferString("")
	for _, v := range keys {
		vs := s[v]
		if len(vs) == 0 {
			b.WriteString(v)
		} else {
			b.WriteString(v + vs[0])
		}
	}

	return b.String()
}

func ValidateSignature(uri, params, signature, authToken string) error {
	u, err := url.ParseQuery(params)
	if err != nil {
		return err
	}

	s := sortedValues(u)
	h := hmac.New(sha1.New, []byte(authToken))

	d := uri + s.String()
	if _, err := h.Write([]byte(d)); err != nil {
		return err
	}
	b := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if b != signature {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
