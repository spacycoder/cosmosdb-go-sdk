package cosmos

func getDummyClient() *Client {
	client, _ := New("AccountEndpoint=https://cosmos-url;AccountKey=abc")
	return client
}
