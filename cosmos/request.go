package cosmos

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HeaderXDate               = "X-Ms-Date"
	HeaderAuth                = "Authorization"
	HeaderVersion             = "X-Ms-Version"
	HeaderContentType         = "Content-Type"
	HeaderContentLength       = "Content-Length"
	HeaderIsQuery             = "X-Ms-Documentdb-Isquery"
	HeaderUpsert              = "x-ms-documentdb-is-upsert"
	HeaderPartitionKey        = "x-ms-documentdb-partitionkey"
	HeaderMaxItemCount        = "x-ms-max-item-count"
	HeaderContinuation        = "x-ms-continuation"
	HeaderConsistency         = "x-ms-consistency-level"
	HeaderSessionToken        = "x-ms-session-token"
	HeaderCrossPartition      = "x-ms-documentdb-query-enablecrosspartition"
	HeaderIfMatch             = "If-Match"
	HeaderIfNonMatch          = "If-None-Match"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderActivityID          = "x-ms-activity-id"
	HeaderRequestCharge       = "x-ms-request-charge"
	HeaderAIM                 = "A-IM"
	HeaderPartitionKeyRangeID = "x-ms-documentdb-partitionkeyrangeid"
	HeaderRetryAfterMs        = "x-ms-retry-after-ms"
	SupportedVersion          = "2018-12-31"
)

type Request struct {
	rLink, rType string
	*http.Request
}

// Return new resource request with type and id
func ResourceRequest(rLink, rType string, req *http.Request) *Request {
	return &Request{rLink, rType, req}
}

// Add 3 default headers to request
// "x-ms-date", "x-ms-version", "authorization"
func (req *Request) DefaultHeaders(mKey string) {
	req.Header.Add(HeaderXDate, formatDate(time.Now()))
	req.Header.Add(HeaderVersion, SupportedVersion)

	tokenToSign := constructTokenString(req.Method, req.rLink, req.rType, req.Header.Get(HeaderXDate))
	token := signAuthToken(tokenToSign, mKey)

	req.Header.Add(HeaderAuth, token)
}

func constructTokenString(method, rLink, rType, date string) string {
	var b strings.Builder
	b.Reset()
	b.WriteString(strings.ToLower(method))
	b.WriteRune('\n')
	b.WriteString(rType)
	b.WriteRune('\n')
	b.WriteString(rLink)
	b.WriteRune('\n')
	b.WriteString(strings.ToLower(date))
	b.WriteRune('\n')
	b.WriteString("")
	b.WriteRune('\n')
	return b.String()
}

func signAuthToken(str string, mkey string) string {
	enc := base64.StdEncoding
	hmacKey, _ := enc.DecodeString(mkey)
	// handle error
	hasher := hmac.New(sha256.New, hmacKey)
	hasher.Write([]byte(str))
	signature := enc.EncodeToString(hasher.Sum(nil))

	authHeader := fmt.Sprintf("type=master&ver=1.0&sig=%s", signature)
	return url.QueryEscape(authHeader)
}

func formatDate(t time.Time) string {
	t = t.UTC()
	return t.Format("Mon, 02 Jan 2006 15:04:05 GMT")
}
