package blockchain

import (
	"net/http"
	"os"
	"testing"
)

func newClient() *Client {
	c := New()
	c.APIKey = os.Getenv("api_key")
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

func TestClient_GetLastError(t *testing.T) {
	c := newClient()
	if c.error != nil {
		t.Fatal("wrong error condition")
	}

	c.setErrorOne(ErrRPE)
	if c.error == nil {
		t.Fatal("wrong error condition")
	}

	e := c.GetLastError()
	if e.ErrMain != ErrRPE || e.Error() != ErrRPE.Error() || c.error != nil {
		t.Fatal("wrong error condition")
	}

	c.setErrorTwo(ErrCRR, ErrCGD)
	e = c.GetLastError()
	if e.ErrMain != ErrCRR || e.ErrExec != ErrCGD {
		t.Fatal("wrong error condition")
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
