package blockchain

import (
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	if NewError(nil, nil, nil, nil) != nil {
		t.Fatal("Wrong error")
	}

	test := "test"
	resp := &http.Response{
		StatusCode: http.StatusOK,
	}
	err := NewError(ErrAIW, ErrBEW, nil, nil)
	if err.Response != nil {
		t.Fatal("wrong Error.Response expected nil: ", *err.Response)
	}
	if err.Address != nil {
		t.Fatal("wrong Error.Address expected nil: ", *err.Address)
	}

	err = NewError(ErrAIW, ErrBEW, resp, &test)
	if err.ErrMain != ErrAIW {
		t.Fatal("wrong Error.ErrMain: ", err.ErrMain)
	}
	if err.ErrExec != ErrBEW {
		t.Fatal("wrong Error.ErrExec", err.ErrExec)
	}
	if *err.Address != test {
		t.Fatal("wrong Error.Address", *err.Address)
	}
	if err.Response.StatusCode != http.StatusOK {
		t.Fatal("wrong Error.Response", *err.Response)
	}
}
