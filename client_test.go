package blockchain

import (
	"net/http"
	"os"
	"testing"
)

const firstBitcoinAddress = "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"

func newClient() *Client {
	c := New()
	c.APIKey = os.Getenv("API_KEY")
	return c
}

func TestNew(t *testing.T) {
	c := New()
	c.BasePath = BasePath
}

func TestNewTor(t *testing.T) {
	c := NewTor()
	c.BasePath = BasePathTor
}

func TestApproveOptions(t *testing.T) {
	var options map[string]string
	options = ApproveOptions(options)
	if options == nil {
		t.Fatal("wrong oprions")
	}

	var s = "test"
	options[s] = s
	options = ApproveOptions(options)
	if options[s] != s {
		t.Fatal("wrong oprions")
	}
}

func TestClient_SetHTTP(t *testing.T) {
	c := New()
	c.SetHTTP(nil)
	if c.client != nil {
		t.Fatal("wrong client condition")
	}

	c.SetHTTP(&http.Client{})
	if c.client == nil {
		t.Fatal("wrong client condition")
	}
}
