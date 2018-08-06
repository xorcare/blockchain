package blockchain

import (
	"os"
	"testing"
)

func newClient() *Client {
	c := New()
	c.ApiKey = os.Getenv("api_key")
	return c
}

func TestNew(t *testing.T) {
	c := New()
	c.BasePath = APIRootNet
}

func TestNewTor(t *testing.T) {
	c := NewTor()
	c.BasePath = APIRootTor
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

	c.setErrorOne(RPE)
	if c.error == nil {
		t.Fatal("wrong error condition")
	}

	e := c.GetLastError()
	if e.ErrorMain != RPE || e.Error() != RPE.Error() || c.error != nil {
		t.Fatal("wrong error condition")
	}

	c.setErrorTwo(CRR, CGD)
	e = c.GetLastError()
	if e.ErrorMain != CRR || e.ErrorExec != CGD {
		t.Fatal("wrong error condition")
	}
}
