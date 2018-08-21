package blockchain

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const firstBitcoinAddress = "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"

func newClient() *Client {
	c := New()
	c.APIKey = os.Getenv("API_KEY")
	c.UserAgent = "tests/" + Version

	return c
}

func customError(e error) error {
	if e == nil {
		return nil
	}

	mess := e.Error()
	if err, ok := e.(*Error); ok {
		mess = ""
		if err.ErrMain != nil {
			mess += fmt.Sprintf("Main error: %s \n", err.ErrMain.Error())
		}
		if err.ErrExec != nil {
			mess += fmt.Sprintf("Execution error: %s \n", err.ErrExec.Error())
		}
		if err.Address != nil {
			mess += fmt.Sprintf("Invalid address is: %34s \n", *err.Address)
		}
		if err.Response != nil {
			mess += fmt.Sprintf("Request url: %s \n", err.Response.Request.URL)
			mess += fmt.Sprintf("Response status (%3d): %s \n", err.Response.StatusCode, err.Response.Status)
			bytes, e := ioutil.ReadAll(err.Response.Body)
			if e == nil {
				mess += fmt.Sprintf("Response body: \n\n %s", string(bytes))
			} else {
				mess += fmt.Sprintf("Response body read error: \n\n %s", e.Error())
			}
		}
	}

	return errors.New(mess)
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
	c.client = nil
	c.SetClient(&http.Client{})
	if c.client == nil {
		t.Fatal("wrong client condition")
	}
}
