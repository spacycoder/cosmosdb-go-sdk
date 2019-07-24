package cosmos

import (
	"net/http"
	"testing"
)

func TestHeaders(t *testing.T) {
	rLink := "dbs/dbtest/colls/colltest/docs/doctest"
	rType := "docs"
	req, err := http.NewRequest("GET", rLink, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	cosmosReq := ResourceRequest(rLink, rType, req)
	cosmosReq.DefaultHeaders("myAccountKey")

	headerDate := cosmosReq.Header.Get(HeaderXDate)
	if headerDate == "" {
		t.Fatal("x-ms-date cannot be empty")
	}

	headerVersion := cosmosReq.Header.Get(HeaderVersion)
	if headerVersion == "" {
		t.Fatal("x-ms-version cannot be empty")
	}
}

func TestTokenCreation(t *testing.T) {
	rLink := "dbs/ToDoList"
	rType := "dbs"
	method := "get"
	date := "Thu, 27 Apr 2017 00:51:12 GMT"
	key := "dsZQi3KtZmCv1ljt3VNWNm7sQUF1y5rJfC6kv5JiwvW0EndXdDku/dkKBp8/ufDToSxLzR4y+O/0H/t4bQtVNw=="
	tokenToSign := constructTokenString(method, rLink, rType, date)
	token := signAuthToken(tokenToSign, key)

	expectedToken := "type%3Dmaster%26ver%3D1.0%26sig%3Dc09PEVJrgp2uQRkr934kFbTqhByc7TVr3OHyqlu%2Bc%2Bc%3D"
	if token != expectedToken {
		t.Fatalf("invalid token expected: %s, got: %s", expectedToken, token)
	}
}
