package cosmos

import "testing"

func getDummyClient() *Client {
	client, _ := New("AccountEndpoint=https://cosmos-url;AccountKey=abc")
	return client
}

func TestEmptyConnString(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Fatal("error should not be nil")
	}
}
